package event

import (
	"context"
	"github.com/golibs-starter/golib/web/event"
	"github.com/khaitq-vnist/auto_ci_be/core/constanst"
	"github.com/khaitq-vnist/auto_ci_be/core/event/message"
)

type UploadFileLogsEvent struct {
	*event.AbstractEvent
	PayloadData *message.LogsEventMessage `json:"payload"`
}

func NewUploadFileLogsEvent(ctx context.Context, eventMessage *message.LogsEventMessage) *UploadFileLogsEvent {
	return &UploadFileLogsEvent{
		AbstractEvent: event.NewAbstractEvent(ctx, constanst.UploadFileLogsEvent),
		PayloadData:   eventMessage,
	}
}
func (a UploadFileLogsEvent) Payload() interface{} {
	return a.PayloadData
}
func (a UploadFileLogsEvent) String() string {
	return a.ToString(a)
}
