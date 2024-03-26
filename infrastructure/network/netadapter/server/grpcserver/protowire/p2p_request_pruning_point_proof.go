package protowire

import (
	"github.com/Kalibriumnet/Kalibrium/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KalibriumdMessage_RequestPruningPointProof) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KalibriumdMessage_RequestPruningPointProof is nil")
	}
	return &appmessage.MsgRequestPruningPointProof{}, nil
}

func (x *KalibriumdMessage_RequestPruningPointProof) fromAppMessage(_ *appmessage.MsgRequestPruningPointProof) error {
	return nil
}
