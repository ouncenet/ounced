package rpchandlers

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/ouncenet/ounced/app/rpc/rpccontext"
	"github.com/ouncenet/ounced/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
