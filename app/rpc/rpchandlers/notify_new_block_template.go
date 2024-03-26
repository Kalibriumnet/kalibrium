package rpchandlers

import (
	"github.com/Kalibriumnet/Kalibrium/app/appmessage"
	"github.com/Kalibriumnet/Kalibrium/app/rpc/rpccontext"
	"github.com/Kalibriumnet/Kalibrium/infrastructure/network/netadapter/router"
)

// HandleNotifyNewBlockTemplate handles the respectively named RPC command
func HandleNotifyNewBlockTemplate(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateNewBlockTemplateNotifications()

	response := appmessage.NewNotifyNewBlockTemplateResponseMessage()
	return response, nil
}
