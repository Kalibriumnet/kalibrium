package blocktemplatebuilder

import (
	"github.com/Kalibriumnet/Kalibrium/domain/consensus/processes/coinbasemanager"
	"github.com/Kalibriumnet/Kalibrium/domain/consensus/utils/merkle"
	"github.com/Kalibriumnet/Kalibrium/domain/consensus/utils/transactionhelper"
	"github.com/Kalibriumnet/Kalibrium/domain/consensusreference"
	"github.com/Kalibriumnet/Kalibrium/util/mstime"
	"math"
	"sort"

	"github.com/Kalibriumnet/Kalibrium/util/difficulty"

	consensusexternalapi "github.com/Kalibriumnet/Kalibrium/domain/consensus/model/externalapi"
	"github.com/Kalibriumnet/Kalibrium/domain/consensus/ruleerrors"
	"github.com/Kalibriumnet/Kalibrium/domain/consensus/utils/subnetworks"
	miningmanagerapi "github.com/Kalibriumnet/Kalibrium/domain/miningmanager/model"
	"github.com/pkg/errors"
)

type candidateTx struct {
	*consensusexternalapi.DomainTransaction
	txValue  float64
	gasLimit uint64

	p     float64
	start float64
	end   float64

	isMarkedForDeletion bool
}

// blockTemplateBuilder creates block templates for a miner to consume
type blockTemplateBuilder struct {
	consensusReference consensusreference.ConsensusReference
	mempool            miningmanagerapi.Mempool
	policy             policy

	coinbasePayloadScriptPublicKeyMaxLength uint8
}

// New creates a new blockTemplateBuilder
func New(consensusReference consensusreference.ConsensusReference, mempool miningmanagerapi.Mempool,
	blockMaxMass uint64, coinbasePayloadScriptPublicKeyMaxLength uint8) miningmanagerapi.BlockTemplateBuilder {
	return &blockTemplateBuilder{
		consensusReference: consensusReference,
		mempool:            mempool,
		policy:             policy{BlockMaxMass: blockMaxMass},

		coinbasePayloadScriptPublicKeyMaxLength: coinbasePayloadScriptPublicKeyMaxLength,
	}
}

// BuildBlockTemplate creates a block template for a miner to consume
// BuildBlockTemplate returns a new block template that is ready to be solved
// using the transactions from the passed transaction source pool and a coinbase
// that either pays to the passed address if it is not nil, or a coinbase that
// is redeemable by anyone if the passed address is nil. The nil address
// functionality is useful since there are cases such as the getblocktemplate
// RPC where external mining software is responsible for creating their own
// coinbase which will replace the one generated for the block template. Thus
// the need to have configured address can be avoided.
//
// The transactions selected and included are prioritized according to several
// factors. First, each transaction has a priority calculated based on its
// value, age of inputs, and size. Transactions which consist of larger
// amounts, older inputs, and small sizes have the highest priority. Second, a
// fee per kilobyte is calculated for each transaction. Transactions with a
// higher fee per kilobyte are preferred. Finally, the block generation related
// policy settings are all taken into account.
//
// Transactions which only spend outputs from other transactions already in the
// block DAG are immediately added to a priority queue which either
// prioritizes based on the priority (then fee per kilobyte) or the fee per
// kilobyte (then priority) depending on whether or not the BlockPrioritySize
// policy setting allots space for high-priority transactions. Transactions
// which spend outputs from other transactions in the source pool are added to a
// dependency map so they can be added to the priority queue once the
// transactions they depend on have been included.
//
// Once the high-priority area (if configured) has been filled with
// transactions, or the priority falls below what is considered high-priority,
// the priority queue is updated to prioritize by fees per kilobyte (then
// priority).
//
// When the fees per kilobyte drop below the TxMinFreeFee policy setting, the
// transaction will be skipped unless the BlockMinSize policy setting is
// nonzero, in which case the block will be filled with the low-fee/free
// transactions until the block size reaches that minimum size.
//
// Any transactions which would cause the block to exceed the BlockMaxMass
// policy setting, exceed the maximum allowed signature operations per block, or
// otherwise cause the block to be invalid are skipped.
//
// Given the above, a block generated by this function is of the following form:
//
//   -----------------------------------  --  --
//  |      Coinbase Transaction         |   |   |
//  |-----------------------------------|   |   |
//  |                                   |   |   | ----- policy.BlockPrioritySize
//  |   High-priority Transactions      |   |   |
//  |                                   |   |   |
//  |-----------------------------------|   | --
//  |                                   |   |
//  |                                   |   |
//  |                                   |   |--- policy.BlockMaxMass
//  |  Transactions prioritized by fee  |   |
//  |  until <= policy.TxMinFreeFee     |   |
//  |                                   |   |
//  |                                   |   |
//  |                                   |   |
//  |-----------------------------------|   |
//  |  Low-fee/Non high-priority (free) |   |
//  |  transactions (while block size   |   |
//  |  <= policy.BlockMinSize)          |   |
//   -----------------------------------  --

func (btb *blockTemplateBuilder) BuildBlockTemplate(
	coinbaseData *consensusexternalapi.DomainCoinbaseData) (*consensusexternalapi.DomainBlockTemplate, error) {

	mempoolTransactions := btb.mempool.BlockCandidateTransactions()
	candidateTxs := make([]*candidateTx, 0, len(mempoolTransactions))
	for _, tx := range mempoolTransactions {
		// Calculate the tx value
		gasLimit := uint64(0)
		if !subnetworks.IsBuiltInOrNative(tx.SubnetworkID) {
			panic("We currently don't support non native subnetworks")
		}
		candidateTxs = append(candidateTxs, &candidateTx{
			DomainTransaction: tx,
			txValue:           btb.calcTxValue(tx),
			gasLimit:          gasLimit,
		})
	}

	// Sort the candidate txs by subnetworkID.
	sort.Slice(candidateTxs, func(i, j int) bool {
		return subnetworks.Less(candidateTxs[i].SubnetworkID, candidateTxs[j].SubnetworkID)
	})

	log.Debugf("Considering %d transactions for inclusion to new block",
		len(candidateTxs))

	blockTxs := btb.selectTransactions(candidateTxs)
	blockTemplate, err := btb.consensusReference.Consensus().BuildBlockTemplate(coinbaseData, blockTxs.selectedTxs)

	invalidTxsErr := ruleerrors.ErrInvalidTransactionsInNewBlock{}
	if errors.As(err, &invalidTxsErr) {
		log.Criticalf("consensusReference.Consensus().BuildBlock returned invalid txs in BuildBlockTemplate")
		err = btb.mempool.RemoveInvalidTransactions(&invalidTxsErr)
		if err != nil {
			// mempool.RemoveInvalidTransactions might return errors in situations that are perfectly fine in this context.
			// TODO: Once the mempool invariants are clear, this should be converted back `return nil, err`:
			// https://github.com/Kalibriumnet/Kalibrium/issues/1553
			log.Criticalf("Error from mempool.RemoveInvalidTransactions: %+v", err)
		}
		// We can call this recursively without worry because this should almost never happen
		return btb.BuildBlockTemplate(coinbaseData)
	}

	if err != nil {
		return nil, err
	}

	log.Debugf("Created new block template (%d transactions, %d in fees, %d mass, target difficulty %064x)",
		len(blockTemplate.Block.Transactions), blockTxs.totalFees, blockTxs.totalMass, difficulty.CompactToBig(blockTemplate.Block.Header.Bits()))

	return blockTemplate, nil
}

// ModifyBlockTemplate modifies an existing block template to the requested coinbase data and updates the timestamp
func (btb *blockTemplateBuilder) ModifyBlockTemplate(newCoinbaseData *consensusexternalapi.DomainCoinbaseData,
	blockTemplateToModify *consensusexternalapi.DomainBlockTemplate) (*consensusexternalapi.DomainBlockTemplate, error) {

	// The first transaction is always the coinbase transaction
	coinbaseTx := blockTemplateToModify.Block.Transactions[transactionhelper.CoinbaseTransactionIndex]
	newPayload, err := coinbasemanager.ModifyCoinbasePayload(coinbaseTx.Payload, newCoinbaseData, btb.coinbasePayloadScriptPublicKeyMaxLength)
	if err != nil {
		return nil, err
	}
	coinbaseTx.Payload = newPayload
	if blockTemplateToModify.CoinbaseHasRedReward {
		// The last output is always the coinbase red blocks reward
		coinbaseTx.Outputs[len(coinbaseTx.Outputs)-1].ScriptPublicKey = newCoinbaseData.ScriptPublicKey
	}
	// Update the hash merkle root according to the modified transactions
	mutableHeader := blockTemplateToModify.Block.Header.ToMutable()
	// TODO: can be optimized to O(log(#transactions)) by caching the whole merkle tree in BlockTemplate and changing only the relevant path
	mutableHeader.SetHashMerkleRoot(merkle.CalculateHashMerkleRoot(blockTemplateToModify.Block.Transactions))

	newTimestamp := mstime.Now().UnixMilliseconds()
	if newTimestamp >= mutableHeader.TimeInMilliseconds() {
		// Only if new time stamp is later than current, update the header. Otherwise,
		// we keep the previous time as built by internal consensus median time logic
		mutableHeader.SetTimeInMilliseconds(newTimestamp)
	}

	blockTemplateToModify.Block.Header = mutableHeader.ToImmutable()
	blockTemplateToModify.CoinbaseData = newCoinbaseData

	return blockTemplateToModify, nil
}

// calcTxValue calculates a value to be used in transaction selection.
// The higher the number the more likely it is that the transaction will be
// included in the block.
func (btb *blockTemplateBuilder) calcTxValue(tx *consensusexternalapi.DomainTransaction) float64 {
	massLimit := btb.policy.BlockMaxMass

	mass := tx.Mass
	fee := tx.Fee
	if subnetworks.IsBuiltInOrNative(tx.SubnetworkID) {
		return float64(fee) / (float64(mass) / float64(massLimit))
	}
	// TODO: Replace with real gas once implemented
	gasLimit := uint64(math.MaxUint64)
	return float64(fee) / (float64(mass)/float64(massLimit) + float64(tx.Gas)/float64(gasLimit))
}
