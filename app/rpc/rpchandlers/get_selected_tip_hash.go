package rpchandlers

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/ouncenet/ounced/app/rpc/rpccontext"
	"github.com/ouncenet/ounced/infrastructure/network/netadapter/router"
)

// HandleGetSelectedTipHash handles the respectively named RPC command
func HandleGetSelectedTipHash(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	selectedTip, err := context.Domain.Consensus().GetVirtualSelectedParent()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetSelectedTipHashResponseMessage(selectedTip.String())

	return response, nil
}
