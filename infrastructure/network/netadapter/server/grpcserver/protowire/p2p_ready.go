package protowire

import (
	"github.com/kalibriumnet/kalibrium/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KalibriumdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KalibriumdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *KalibriumdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
