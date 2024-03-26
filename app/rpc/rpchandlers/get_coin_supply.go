package rpchandlers

import (
	"github.com/Kalibriumnet/Kalibrium/app/appmessage"
	"github.com/Kalibriumnet/Kalibrium/app/rpc/rpccontext"
	"github.com/Kalibriumnet/Kalibrium/domain/consensus/utils/constants"
	"github.com/Kalibriumnet/Kalibrium/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when Kalibrium is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingEquilSupply, err := context.UTXOIndex.GetCirculatingEquilSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxEquil,
		circulatingEquilSupply,
	)

	return response, nil
}
