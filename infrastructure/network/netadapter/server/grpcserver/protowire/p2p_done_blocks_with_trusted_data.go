package protowire

import (
	"github.com/Kalibriumnet/Kalibrium/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KalibriumdMessage_DoneBlocksWithTrustedData) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KalibriumdMessage_DoneBlocksWithTrustedData is nil")
	}
	return &appmessage.MsgDoneBlocksWithTrustedData{}, nil
}

func (x *KalibriumdMessage_DoneBlocksWithTrustedData) fromAppMessage(_ *appmessage.MsgDoneBlocksWithTrustedData) error {
	return nil
}
