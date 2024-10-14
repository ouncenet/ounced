package protowire

import (
	"github.com/ouncenet/ounced/app/appmessage"
	"github.com/pkg/errors"
)

func (x *OuncedMessage_NotifyFinalityConflictsRequest) toAppMessage() (appmessage.Message, error) {
	return &appmessage.NotifyFinalityConflictsRequestMessage{}, nil
}

func (x *OuncedMessage_NotifyFinalityConflictsRequest) fromAppMessage(_ *appmessage.NotifyFinalityConflictsRequestMessage) error {
	x.NotifyFinalityConflictsRequest = &NotifyFinalityConflictsRequestMessage{}
	return nil
}

func (x *OuncedMessage_NotifyFinalityConflictsResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_NotifyFinalityConflictsResponse is nil")
	}
	return x.NotifyFinalityConflictsResponse.toAppMessage()
}

func (x *OuncedMessage_NotifyFinalityConflictsResponse) fromAppMessage(message *appmessage.NotifyFinalityConflictsResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.NotifyFinalityConflictsResponse = &NotifyFinalityConflictsResponseMessage{
		Error: err,
	}
	return nil
}

func (x *NotifyFinalityConflictsResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "NotifyFinalityConflictsResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.NotifyFinalityConflictsResponseMessage{
		Error: rpcErr,
	}, nil
}

func (x *OuncedMessage_FinalityConflictNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_FinalityConflictNotification is nil")
	}
	return x.FinalityConflictNotification.toAppMessage()
}

func (x *OuncedMessage_FinalityConflictNotification) fromAppMessage(message *appmessage.FinalityConflictNotificationMessage) error {
	x.FinalityConflictNotification = &FinalityConflictNotificationMessage{
		ViolatingBlockHash: message.ViolatingBlockHash,
	}
	return nil
}

func (x *FinalityConflictNotificationMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "FinalityConflictNotificationMessage is nil")
	}
	return &appmessage.FinalityConflictNotificationMessage{
		ViolatingBlockHash: x.ViolatingBlockHash,
	}, nil
}

func (x *OuncedMessage_FinalityConflictResolvedNotification) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "OuncedMessage_FinalityConflictResolvedNotification is nil")
	}
	return x.FinalityConflictResolvedNotification.toAppMessage()
}

func (x *OuncedMessage_FinalityConflictResolvedNotification) fromAppMessage(message *appmessage.FinalityConflictResolvedNotificationMessage) error {
	x.FinalityConflictResolvedNotification = &FinalityConflictResolvedNotificationMessage{
		FinalityBlockHash: message.FinalityBlockHash,
	}
	return nil
}

func (x *FinalityConflictResolvedNotificationMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "FinalityConflictResolvedNotificationMessage is nil")
	}
	return &appmessage.FinalityConflictResolvedNotificationMessage{
		FinalityBlockHash: x.FinalityBlockHash,
	}, nil
}
