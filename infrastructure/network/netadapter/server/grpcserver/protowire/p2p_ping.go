package protowire

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/pkg/errors"
)

func (x *OuncedMessage_Ping) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_Ping is nil")
	}
	return x.Ping.toAppMessage()
}

func (x *PingMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PingMessage is nil")
	}
	return &appmessage.MsgPing{
		Nonce: x.Nonce,
	}, nil
}

func (x *OuncedMessage_Ping) fromAppMessage(msgPing *appmessage.MsgPing) error {
	x.Ping = &PingMessage{
		Nonce: msgPing.Nonce,
	}
	return nil
}
