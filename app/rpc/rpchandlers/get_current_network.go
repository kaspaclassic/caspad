package rpchandlers

import (
	"github.com/kaspaclassic/caspad/app/appmessage"
	"github.com/kaspaclassic/caspad/app/rpc/rpccontext"
	"github.com/kaspaclassic/caspad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
