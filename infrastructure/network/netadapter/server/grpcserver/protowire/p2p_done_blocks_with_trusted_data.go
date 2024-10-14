package protowire

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/pkg/errors"
)

func (x *OuncedMessage_DoneBlocksWithTrustedData) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_DoneBlocksWithTrustedData is nil")
	}
	return &appmessage.MsgDoneBlocksWithTrustedData{}, nil
}

func (x *OuncedMessage_DoneBlocksWithTrustedData) fromAppMessage(_ *appmessage.MsgDoneBlocksWithTrustedData) error {
	return nil
}
