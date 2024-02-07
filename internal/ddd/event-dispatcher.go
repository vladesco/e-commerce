package ddd

import (
	"context"
	"sync"
)

type (
	EventHandler[T Event] interface {
		HandleEvent(ctx context.Context, event T) error
	}

	EventHandlerFunc[T Event] func(ctx context.Context, event T) error

	EventSubscriber[T Event] interface {
		Subscribe(eventName string, handler EventHandler[T])
	}

	EventPublisher[T Event] interface {
		Publish(ctx context.Context, events ...T) error
	}

	EventDispatcher[T Event] struct {
		mutex    sync.Mutex
		handlers map[string][]EventHandler[T]
	}
)

func NewEventDispatcher[T Event]() *EventDispatcher[T] {
	return &EventDispatcher[T]{
		handlers: make(map[string][]EventHandler[T]),
	}
}

func (dispatcher *EventDispatcher[T]) Subscribe(eventName string, handler EventHandler[T]) {
	dispatcher.mutex.Lock()
	defer dispatcher.mutex.Unlock()

	dispatcher.handlers[eventName] = append(dispatcher.handlers[eventName], handler)
}

func (dispatcher *EventDispatcher[T]) Publish(ctx context.Context, events ...T) error {
	for _, event := range events {
		eventName := event.GetName()

		for _, handler := range dispatcher.handlers[eventName] {
			err := handler.HandleEvent(ctx, event)

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (function EventHandlerFunc[T]) HandleEvent(ctx context.Context, event T) error {
	return function(ctx, event)
}

// CHECKS
var _ interface {
	EventSubscriber[Event]
	EventPublisher[Event]
} = (*EventDispatcher[Event])(nil)
