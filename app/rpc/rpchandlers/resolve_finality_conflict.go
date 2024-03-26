package rpchandlers

import (
	"github.com/casklas/caspad/app/appmessage"
	"github.com/casklas/caspad/app/rpc/rpccontext"
	"github.com/casklas/caspad/infrastructure/network/netadapter/router"
)

// HandleResolveFinalityConflict handles the respectively named RPC command
func HandleResolveFinalityConflict(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	if context.Config.SafeRPC {
		log.Warn("ResolveFinalityConflict RPC command called while node in safe RPC mode -- ignoring.")
		response := &appmessage.ResolveFinalityConflictResponseMessage{}
		response.Error =
			appmessage.RPCErrorf("ResolveFinalityConflict RPC command called while node in safe RPC mode")
		return response, nil
	}

	response := &appmessage.ResolveFinalityConflictResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
