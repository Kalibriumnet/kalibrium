package rpchandlers

import (
	"github.com/kalibriumnet/kalibrium/app/appmessage"
	"github.com/kalibriumnet/kalibrium/app/rpc/rpccontext"
	"github.com/kalibriumnet/kalibrium/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
