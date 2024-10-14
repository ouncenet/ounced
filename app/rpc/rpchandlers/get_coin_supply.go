package rpchandlers

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/ouncenet/ounced/app/rpc/rpccontext"
	"github.com/ouncenet/ounced/domain/consensus/utils/constants"
	"github.com/ouncenet/ounced/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when kaspad is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingGrainSupply, err := context.UTXOIndex.GetCirculatingGrainSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxGrain,
		circulatingGrainSupply,
	)

	return response, nil
}
