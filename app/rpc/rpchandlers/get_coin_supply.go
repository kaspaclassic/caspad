package rpchandlers

import (
	"github.com/casklas/caspad/app/appmessage"
	"github.com/casklas/caspad/app/rpc/rpccontext"
	"github.com/casklas/caspad/domain/consensus/utils/constants"
	"github.com/casklas/caspad/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when caspad is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingLeorSupply, err := context.UTXOIndex.GetCirculatingLeorSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxLeor,
		circulatingLeorSupply,
	)

	return response, nil
}
