package testapi

import (
	"github.com/kalibriumnet/kalibrium/domain/consensus/model"
	"github.com/kalibriumnet/kalibrium/domain/consensus/utils/txscript"
)

// TestTransactionValidator adds to the main TransactionValidator methods required by tests
type TestTransactionValidator interface {
	model.TransactionValidator
	SigCache() *txscript.SigCache
	SetSigCache(sigCache *txscript.SigCache)
}
