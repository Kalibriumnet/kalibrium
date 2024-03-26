package protowire

import "github.com/Kalibriumnet/Kalibrium/app/appmessage"

func (x *KalibriumdMessage_UnexpectedPruningPoint) toAppMessage() (appmessage.Message, error) {
	return &appmessage.MsgUnexpectedPruningPoint{}, nil
}

func (x *KalibriumdMessage_UnexpectedPruningPoint) fromAppMessage(_ *appmessage.MsgUnexpectedPruningPoint) error {
	return nil
}
