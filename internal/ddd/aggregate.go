package ddd

const (
	AggregateIDKey      = "aggregate-id"
	AggregateNameKey    = "aggregate-name"
	AggregateVersionKey = "aggregate-version"
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
		GetAggregateVersion() int
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
			AggregateNameKey: aggregate.name,
			AggregateIDKey:   aggregate.id,
		})

	aggregate.events = append(
		aggregate.events,
		&aggregateEvent{
			newEvent(eventName, payload, options...),
		},
	)
}

func (event *aggregateEvent) GetAggregateId() string {
	return event.metadata.Get(AggregateIDKey).(string)
}
func (event *aggregateEvent) GetAggregateName() string {
	return event.metadata.Get(AggregateNameKey).(string)
}
func (event *aggregateEvent) GetAggregateVersion() int {
	return event.metadata.Get(AggregateVersionKey).(int)
}

// CHECKS
var _ interface {
	Ider
	Namer
	Eventer[AggregateEvent]
} = (*Aggregate)(nil)

var _ AggregateEvent = (*aggregateEvent)(nil)
