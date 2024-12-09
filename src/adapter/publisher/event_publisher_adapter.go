package publisher

import (
	"github.com/golibs-starter/golib/pubsub"
	"github.com/khaitq-vnist/auto_ci_be/core/port"
)

type EventPublisherAdapter struct {
}

func (e2 EventPublisherAdapter) Publish(e pubsub.Event) {
	pubsub.Publish(e)
}

func NewEventPublisherAdapter() port.IEventPublisherPort {
	return &EventPublisherAdapter{}
}
