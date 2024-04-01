package protowire

import (
	"github.com/kalibriumnet/kalibrium/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KalibriumdMessage_RequestIBDBlocks) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KalibriumdMessage_RequestIBDBlocks is nil")
	}
	return x.RequestIBDBlocks.toAppMessage()
}

func (x *RequestIBDBlocksMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RequestIBDBlocksMessage is nil")
	}
	hashes, err := protoHashesToDomain(x.Hashes)
	if err != nil {
		return nil, err
	}
	return &appmessage.MsgRequestIBDBlocks{Hashes: hashes}, nil
}

func (x *KalibriumdMessage_RequestIBDBlocks) fromAppMessage(msgRequestIBDBlocks *appmessage.MsgRequestIBDBlocks) error {
	x.RequestIBDBlocks = &RequestIBDBlocksMessage{
		Hashes: domainHashesToProto(msgRequestIBDBlocks.Hashes),
	}

	return nil
}
