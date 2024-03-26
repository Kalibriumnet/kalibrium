package libkalibriumwallet

import (
	"encoding/hex"

	"github.com/Kalibriumnet/Kalibrium/app/appmessage"
	"github.com/Kalibriumnet/Kalibrium/cmd/kalibriumwallet/daemon/pb"
	"github.com/Kalibriumnet/Kalibrium/domain/consensus/model/externalapi"
	"github.com/Kalibriumnet/Kalibrium/domain/consensus/utils/transactionid"
	"github.com/Kalibriumnet/Kalibrium/domain/consensus/utils/utxo"
)

// KalibriumwalletdUTXOsTolibkalibriumwalletUTXOs converts a  []*pb.UtxosByAddressesEntry to a []*libkalibriumwallet.UTXO
func KalibriumwalletdUTXOsTolibkalibriumwalletUTXOs(kalibriumwalletdUtxoEntires []*pb.UtxosByAddressesEntry) ([]*UTXO, error) {
	UTXOs := make([]*UTXO, len(kalibriumwalletdUtxoEntires))
	for i, entry := range kalibriumwalletdUtxoEntires {
		script, err := hex.DecodeString(entry.UtxoEntry.ScriptPublicKey.ScriptPublicKey)
		if err != nil {
			return nil, err
		}
		transactionID, err := transactionid.FromString(entry.Outpoint.TransactionId)
		if err != nil {
			return nil, err
		}
		UTXOs[i] = &UTXO{
			UTXOEntry: utxo.NewUTXOEntry(
				entry.UtxoEntry.Amount,
				&externalapi.ScriptPublicKey{
					Script:  script,
					Version: uint16(entry.UtxoEntry.ScriptPublicKey.Version),
				},
				entry.UtxoEntry.IsCoinbase,
				entry.UtxoEntry.BlockDaaScore,
			),
			Outpoint: &externalapi.DomainOutpoint{
				TransactionID: *transactionID,
				Index:         entry.Outpoint.Index,
			},
		}
	}
	return UTXOs, nil
}

// AppMessageUTXOToKalibriumwalletdUTXO converts an appmessage.UTXOsByAddressesEntry to a  pb.UtxosByAddressesEntry
func AppMessageUTXOToKalibriumwalletdUTXO(appUTXOsByAddressesEntry *appmessage.UTXOsByAddressesEntry) *pb.UtxosByAddressesEntry {
	return &pb.UtxosByAddressesEntry{
		Outpoint: &pb.Outpoint{
			TransactionId: appUTXOsByAddressesEntry.Outpoint.TransactionID,
			Index:         appUTXOsByAddressesEntry.Outpoint.Index,
		},
		UtxoEntry: &pb.UtxoEntry{
			Amount: appUTXOsByAddressesEntry.UTXOEntry.Amount,
			ScriptPublicKey: &pb.ScriptPublicKey{
				Version:         uint32(appUTXOsByAddressesEntry.UTXOEntry.ScriptPublicKey.Version),
				ScriptPublicKey: appUTXOsByAddressesEntry.UTXOEntry.ScriptPublicKey.Script,
			},
			BlockDaaScore: appUTXOsByAddressesEntry.UTXOEntry.BlockDAAScore,
			IsCoinbase:    appUTXOsByAddressesEntry.UTXOEntry.IsCoinbase,
		},
	}
}
