package protowire

import (
	"github.com/Kalibriumnet/Kalibrium/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KalibriumdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KalibriumdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *KalibriumdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
