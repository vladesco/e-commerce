package ddd

import "github.com/google/uuid"

type Aggregate interface {
	Entity
	AddEvent(event Event)
	GetEvents() []Event
}

type AggregateBase struct {
	id     string
	events []Event
}

func (aggregate *AggregateBase) GetId() string {
	return aggregate.id
}

func (aggregate *AggregateBase) AddEvent(event Event) {
	aggregate.events = append(aggregate.events, event)
}

func (aggregate *AggregateBase) GetEvents() []Event {
	return aggregate.events
}

func CrateAggregate() AggregateBase {
	return AggregateBase{
		id:     uuid.NewString(),
		events: []Event{},
	}
}
