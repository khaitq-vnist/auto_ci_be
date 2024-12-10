package usecase

import (
	"context"
	event2 "github.com/khaitq-vnist/auto_ci_be/core/event"
	message2 "github.com/khaitq-vnist/auto_ci_be/core/event/message"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type IFireEventUsecase interface {
	PublishUploadLogsEvent(ctx context.Context, payload *message2.LogsEventMessage)
}
type FireEventUsecase struct {
	eventPublisherPort port.IEventPublisherPort
}

func (f FireEventUsecase) PublishUploadLogsEvent(ctx context.Context, payload *message2.LogsEventMessage) {
	event := event2.NewUploadFileLogsEvent(ctx, payload)
	f.eventPublisherPort.Publish(event)
}

func NewFireEventUsecase(eventPublisherPort port.IEventPublisherPort) IFireEventUsecase {
	return &FireEventUsecase{
		eventPublisherPort: eventPublisherPort,
	}
}
