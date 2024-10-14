package protowire

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/pkg/errors"
)

func (x *OuncedMessage_RequestNextHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_RequestNextHeaders is nil")
	}
	return &appmessage.MsgRequestNextHeaders{}, nil
}

func (x *OuncedMessage_RequestNextHeaders) fromAppMessage(_ *appmessage.MsgRequestNextHeaders) error {
	return nil
}
