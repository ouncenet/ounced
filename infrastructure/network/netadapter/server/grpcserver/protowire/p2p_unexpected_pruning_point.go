package protowire

import "github.com/ouncenet/ounced/app/appmessage"

func (x *OuncedMessage_UnexpectedPruningPoint) toAppMessage() (appmessage.Message, error) {
	return &appmessage.MsgUnexpectedPruningPoint{}, nil
}

func (x *OuncedMessage_UnexpectedPruningPoint) fromAppMessage(_ *appmessage.MsgUnexpectedPruningPoint) error {
	return nil
}
