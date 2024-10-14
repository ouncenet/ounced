package protowire

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/pkg/errors"
)

func (x *OuncedMessage_RequestPruningPointProof) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_RequestPruningPointProof is nil")
	}
	return &appmessage.MsgRequestPruningPointProof{}, nil
}

func (x *OuncedMessage_RequestPruningPointProof) fromAppMessage(_ *appmessage.MsgRequestPruningPointProof) error {
	return nil
}
