package eventsourcing

import (
	"context"

	"github.com/vladesco/e-commerce/internal/ddd"
)

type (
	EventSourcedAggregate interface {
		ddd.Ider
		ddd.Namer
		ddd.Eventer[AggregateEvent]
		Versioner
		EventApplier
		EventCommitter
	}

	AggregateStoreMiddleware func(store AggregateStore) AggregateStore

	AggregateStore interface {
		Load(ctx context.Context, aggregate EventSourcedAggregate) error
		Save(ctx context.Context, aggregate EventSourcedAggregate) error
	}
)

func AggregateStoreWithMiddleware(store AggregateStore, middlewares ...AggregateStoreMiddleware) AggregateStore {
	for _, middelware := range middlewares {
		store = middelware(store)
	}

	return store
}
