package protowire

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/pkg/errors"
)

func (x *OuncedMessage_RequestNextPruningPointAndItsAnticoneBlocks) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_DonePruningPointAndItsAnticoneBlocks is nil")
	}
	return &appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks{}, nil
}

func (x *OuncedMessage_RequestNextPruningPointAndItsAnticoneBlocks) fromAppMessage(_ *appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks) error {
	return nil
}
