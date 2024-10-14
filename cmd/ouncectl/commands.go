package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/ouncenet/ounced/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.OuncedMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.OuncedMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.OuncedMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.OuncedMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.OuncedMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.OuncedMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.OuncedMessage_BanRequest{}),
	reflect.TypeOf(protowire.OuncedMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
