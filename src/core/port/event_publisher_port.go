package port

import "github.com/golibs-starter/golib/pubsub"

type IEventPublisherPort interface {
	Publish(e pubsub.Event)
}
