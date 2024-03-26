package protowire

import (
	"github.com/Kalibriumnet/Kalibrium/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KalibriumdMessage_IbdBlock) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KalibriumdMessage_IbdBlock is nil")
	}
	msgBlock, err := x.IbdBlock.toAppMessage()
	if err != nil {
		return nil, err
	}
	return &appmessage.MsgIBDBlock{MsgBlock: msgBlock}, nil
}

func (x *KalibriumdMessage_IbdBlock) fromAppMessage(msgIBDBlock *appmessage.MsgIBDBlock) error {
	x.IbdBlock = new(BlockMessage)
	return x.IbdBlock.fromAppMessage(msgIBDBlock.MsgBlock)
}
