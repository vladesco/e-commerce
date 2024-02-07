package ddd

const (
	aggregateIdKey   = "aggregate-id"
	aggregateNameKey = "aggregate-name"
)

type (
	Aggregate struct {
		Entity
		events []AggregateEvent
	}

	AggregateEvent interface {
		Event
		GetAggregateId() string
		GetAggregateName() string
	}

	aggregateEvent struct {
		event
	}
)

func NewAggregate(id, name string) Aggregate {
	return Aggregate{
		Entity: NewEntity(id, name),
		events: []AggregateEvent{},
	}
}

func (aggregate *Aggregate) GetEvents() []AggregateEvent       { return aggregate.events }
func (aggregate *Aggregate) ClearEvents()                      { aggregate.events = []AggregateEvent{} }
func (aggregate *Aggregate) SetEvents(events []AggregateEvent) { aggregate.events = events }

func (aggregate *Aggregate) AddEvent(eventName string, payload EventPayload, options ...EventOption) {
	options = append(options,
		Metadata{
			aggregateIdKey:   aggregate.id,
			aggregateNameKey: aggregate.name,
		})

	aggregate.events = append(
		aggregate.events,
		&aggregateEvent{
			newEvent(eventName, payload, options...),
		},
	)
}

func (event *aggregateEvent) GetAggregateId() string {
	return event.metadata.Get(aggregateIdKey).(string)
}
func (event *aggregateEvent) GetAggregateName() string {
	return event.metadata.Get(aggregateNameKey).(string)
}

// CHECKS
var _ interface {
	Ider
	Namer
	Eventer[AggregateEvent]
} = (*Aggregate)(nil)

var _ AggregateEvent = (*aggregateEvent)(nil)
