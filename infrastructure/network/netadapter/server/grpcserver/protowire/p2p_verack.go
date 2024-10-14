package protowire

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/pkg/errors"
)

func (x *OuncedMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *OuncedMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
