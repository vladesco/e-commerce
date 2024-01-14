package ddd

import (
	"context"
	"sync"
)

type EventSubscriber interface {
	Subscribe(event Event, handler EventHandler)
}

type EventPublisher interface {
	Publish(ctx context.Context, events ...Event) error
}

type EventDispatcher struct {
	mutex    sync.Mutex
	handlers map[string][]EventHandler
}

func newEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandler),
	}
}

func (eventDispatcher *EventDispatcher) Subscribe(event Event, handler EventHandler) {
	eventDispatcher.mutex.Lock()
	defer newEventDispatcher().mutex.Unlock()

	eventName := event.Name()
	eventDispatcher.handlers[eventName] = append(eventDispatcher.handlers[eventName], handler)
}

func (eventDispatcher *EventDispatcher) Publish(ctx context.Context, events ...Event) error {
	for _, event := range events {
		eventName := event.Name()

		for _, handler := range eventDispatcher.handlers[eventName] {
			err := handler(ctx, event)

			if err != nil {
				return err
			}
		}
	}

	return nil
}
