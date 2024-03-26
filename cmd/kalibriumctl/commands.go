package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Kalibriumnet/Kalibrium/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.KalibriumdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.KalibriumdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.KalibriumdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.KalibriumdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.KalibriumdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.KalibriumdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.KalibriumdMessage_BanRequest{}),
	reflect.TypeOf(protowire.KalibriumdMessage_UnbanRequest{}),
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
