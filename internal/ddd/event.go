package ddd

import (
	"time"

	"github.com/google/uuid"
)

type (
	EventPayload = any

	Event interface {
		Ider
		Namer
		GetPayload() EventPayload
		GetMetadata() Metadata
		GetOccurredAt() time.Time
	}

	event struct {
		Entity
		payload    EventPayload
		metadata   Metadata
		occurredAt time.Time
	}

	Eventer[T Event] interface {
		AddEvent(string, EventPayload, ...EventOption)
		GetEvents() []T
		ClearEvents()
	}
)

func newEvent(eventName string, payload EventPayload, options ...EventOption) event {
	event := event{
		Entity:     NewEntity(uuid.New().String(), eventName),
		payload:    payload,
		metadata:   make(Metadata),
		occurredAt: time.Now(),
	}

	for _, option := range options {
		option.configureEvent(&event)
	}

	return event
}

func (event *event) GetPayload() EventPayload { return event.payload }
func (event *event) GetMetadata() Metadata    { return event.metadata }
func (event *event) GetOccurredAt() time.Time { return event.occurredAt }

// CHECKS
var _ Event = (*event)(nil)
