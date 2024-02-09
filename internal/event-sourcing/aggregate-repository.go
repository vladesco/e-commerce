package eventsourcing

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/internal/ddd"
	"github.com/vladesco/e-commerce/internal/registarable"
)

type AggregateRepository[T EventSourcedAggregate] struct {
	aggregateName string
	registry      registarable.Registry
	store         AggregateStore
}

func NewAggregateRepository[T EventSourcedAggregate](aggregateName string, registry registarable.Registry, store AggregateStore) *AggregateRepository[T] {
	return &AggregateRepository[T]{
		aggregateName: aggregateName,
		registry:      registry,
		store:         store,
	}
}

func (repository *AggregateRepository[T]) Load(ctx context.Context, aggregateId string) (T, error) {
	var aggregate T
	var isValueConveribleToAggregate bool

	value, err := repository.registry.Build(repository.aggregateName, ddd.SetId(aggregateId), ddd.SetName(repository.aggregateName))

	if err != nil {
		return aggregate, errors.Wrap(err, "error while building aggregate")
	}

	if aggregate, isValueConveribleToAggregate = value.(T); !isValueConveribleToAggregate {
		return aggregate, errors.Wrapf(err, "error while casting %T to aggregate type %T", value, aggregate)
	}

	err = repository.store.Load(ctx, aggregate)

	if err != nil {
		return aggregate, errors.Wrap(err, "error while loading aggregate")
	}

	return aggregate, nil
}

func (repository *AggregateRepository[T]) Save(ctx context.Context, aggregate T) error {
	if aggregate.GetVersion() == aggregate.GetPendingVersion() {
		return nil
	}

	for _, event := range aggregate.GetEvents() {
		if err := aggregate.ApplyEvent(event); err != nil {
			return errors.Wrapf(err, "error while trying apply events for aggregate %T", aggregate)
		}
	}

	err := repository.store.Save(ctx, aggregate)

	if err != nil {
		return errors.Wrapf(err, "error while trying save aggregate %T", aggregate)
	}

	aggregate.CommitEvents()

	return nil
}
