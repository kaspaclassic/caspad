package rpchandlers

import (
	"github.com/casklas/caspad/app/appmessage"
	"github.com/casklas/caspad/app/rpc/rpccontext"
	"github.com/casklas/caspad/infrastructure/network/netadapter/router"
)

// HandleNotifyFinalityConflicts handles the respectively named RPC command
func HandleNotifyFinalityConflicts(context *rpccontext.Context, router *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	listener, err := context.NotificationManager.Listener(router)
	if err != nil {
		return nil, err
	}
	listener.PropagateFinalityConflictNotifications()
	listener.PropagateFinalityConflictResolvedNotifications()

	response := appmessage.NewNotifyFinalityConflictsResponseMessage()
	return response, nil
}
