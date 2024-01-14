package ddd

import "context"

type Event interface {
	Name() string
}

type EventHandler func(ctx context.Context, event Event) error
