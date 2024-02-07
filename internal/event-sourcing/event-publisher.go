package eventsourcing

import (
	"context"

	"github.com/vladesco/e-commerce/internal/ddd"
)

type EventPublisher struct {
	AggregateStore
	publisher ddd.EventPublisher[AggregateEvent]
}

func NewEventPublisherMiddleware(publisher ddd.EventPublisher[AggregateEvent]) AggregateStoreMiddleware {
	eventPublisher := EventPublisher{
		publisher: publisher,
	}

	return func(store AggregateStore) AggregateStore {
		eventPublisher.AggregateStore = store
		return eventPublisher
	}
}

func (eventPublisher EventPublisher) Save(ctx context.Context, aggregate EventSourcedAggregate) error {
	if err := eventPublisher.AggregateStore.Save(ctx, aggregate); err != nil {
		return err
	}
	return eventPublisher.publisher.Publish(ctx, aggregate.GetEvents()...)
}
