package eventsourcing

import (
	"github.com/vladesco/e-commerce/internal/ddd"
)

const aggregateVersionKey = "aggregate-version"

type (
	EventApplier interface {
		ApplyEvent(ddd.Event) error
	}

	EventCommitter interface {
		CommitEvents()
	}

	AggregateEvent interface {
		ddd.AggregateEvent
		GetAggregateVersion() int
	}

	aggregateEvent struct {
		ddd.AggregateEvent
	}
)

func (event *aggregateEvent) GetAggregateVersion() int {
	return event.GetMetadata().Get(aggregateVersionKey).(int)
}

// CHECKS
var _ AggregateEvent = (*aggregateEvent)(nil)
