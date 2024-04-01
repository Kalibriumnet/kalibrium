package protowire

import (
	"github.com/kalibriumnet/kalibrium/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KalibriumdMessage_GetBlockCountRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KalibriumdMessage_GetBlockCountRequest is nil")
	}
	return &appmessage.GetBlockCountRequestMessage{}, nil
}

func (x *KalibriumdMessage_GetBlockCountRequest) fromAppMessage(_ *appmessage.GetBlockCountRequestMessage) error {
	x.GetBlockCountRequest = &GetBlockCountRequestMessage{}
	return nil
}

func (x *KalibriumdMessage_GetBlockCountResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KalibriumdMessage_GetBlockCountResponse is nil")
	}
	return x.GetBlockCountResponse.toAppMessage()
}

func (x *KalibriumdMessage_GetBlockCountResponse) fromAppMessage(message *appmessage.GetBlockCountResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.GetBlockCountResponse = &GetBlockCountResponseMessage{
		BlockCount:  message.BlockCount,
		HeaderCount: message.HeaderCount,
		Error:       err,
	}
	return nil
}

func (x *GetBlockCountResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetBlockCountResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	if rpcErr != nil && (x.BlockCount != 0 || x.HeaderCount != 0) {
		return nil, errors.New("GetBlockCountResponseMessage contains both an error and a response")
	}
	return &appmessage.GetBlockCountResponseMessage{
		BlockCount:  x.BlockCount,
		HeaderCount: x.HeaderCount,
		Error:       rpcErr,
	}, nil
}
