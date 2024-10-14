package protowire

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/pkg/errors"
)

func (x *OuncedMessage_DoneHeaders) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_DoneHeaders is nil")
	}
	return &appmessage.MsgDoneHeaders{}, nil
}

func (x *OuncedMessage_DoneHeaders) fromAppMessage(_ *appmessage.MsgDoneHeaders) error {
	return nil
}
