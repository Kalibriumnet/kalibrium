package rpchandlers

import (
	"github.com/Kalibriumnet/Kalibrium/app/appmessage"
	"github.com/Kalibriumnet/Kalibrium/app/rpc/rpccontext"
	"github.com/Kalibriumnet/Kalibrium/infrastructure/network/netadapter/router"
)

// HandleNotifyBlockAdded handles the respectively named RPC command
func HandleNotifyBlockAdded(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateBlockAddedNotifications()

	response := appmessage.NewNotifyBlockAddedResponseMessage()
	return response, nil
}
