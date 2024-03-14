package rpchandlers

import (
	"github.com/casklas/caspad/app/appmessage"
	"github.com/casklas/caspad/app/rpc/rpccontext"
	"github.com/casklas/caspad/domain/consensus/model/externalapi"
	"github.com/casklas/caspad/domain/consensus/utils/transactionhelper"
	"github.com/casklas/caspad/domain/consensus/utils/txscript"
	"github.com/casklas/caspad/infrastructure/network/netadapter/router"
	"github.com/casklas/caspad/util"
	"github.com/casklas/caspad/version"
)

// HandleGetBlockTemplate handles the respectively named RPC command
func HandleGetBlockTemplate(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	getBlockTemplateRequest := request.(*appmessage.GetBlockTemplateRequestMessage)

	payAddress, err := util.DecodeAddress(getBlockTemplateRequest.PayAddress, context.Config.ActiveNetParams.Prefix)
	if err != nil {
		errorMessage := &appmessage.GetBlockTemplateResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Could not decode address: %s", err)
		return errorMessage, nil
	}

	scriptPublicKey, err := txscript.PayToAddrScript(payAddress)
	if err != nil {
		return nil, err
	}

	coinbaseData := &externalapi.DomainCoinbaseData{ScriptPublicKey: scriptPublicKey, ExtraData: []byte(version.Version() + "/" + getBlockTemplateRequest.ExtraData)}

	templateBlock, isNearlySynced, err := context.Domain.MiningManager().GetBlockTemplate(coinbaseData)
	if err != nil {
		return nil, err
	}

	if uint64(len(templateBlock.Transactions[transactionhelper.CoinbaseTransactionIndex].Payload)) > context.Config.NetParams().MaxCoinbasePayloadLength {
		errorMessage := &appmessage.GetBlockTemplateResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Coinbase payload is above max length (%d). Try to shorten the extra data.", context.Config.NetParams().MaxCoinbasePayloadLength)
		return errorMessage, nil
	}

	rpcBlock := appmessage.DomainBlockToRPCBlock(templateBlock)

	return appmessage.NewGetBlockTemplateResponseMessage(rpcBlock, context.ProtocolManager.Context().HasPeers() && isNearlySynced), nil
}
