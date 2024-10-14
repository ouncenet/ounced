package protowire

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/pkg/errors"
)

func (x *OuncedMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *OuncedMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
