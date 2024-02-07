package eventsourcing

import "github.com/vladesco/e-commerce/internal/ddd"

type Aggregate struct {
	ddd.Aggregate
	version int
}

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Aggregate: ddd.NewAggregate(id, name),
		version:   0,
	}
}

func (aggregate *Aggregate) AddEvent(eventName string, payload ddd.EventPayload, options ...ddd.EventOption) {
	options = append(options, ddd.Metadata{
		aggregateVersionKey: aggregate.GetPendingVersion() + 1,
	})

	aggregate.Aggregate.AddEvent(eventName, payload, options...)
}

func (aggregate *Aggregate) CommitEvents() {
	aggregate.version += aggregate.GetPendingVersion()
	aggregate.ClearEvents()
}

func (aggregate *Aggregate) GetVersion() int { return aggregate.version }

func (aggregate *Aggregate) GetPendingVersion() int {
	return aggregate.version + len(aggregate.GetEvents())
}

func (aggregate *Aggregate) setVersion(version int) {
	aggregate.version = version
}

// CHECKS
var _ interface {
	Versioner
	EventCommitter
} = (*Aggregate)(nil)
