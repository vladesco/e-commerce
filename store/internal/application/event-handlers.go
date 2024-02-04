package application

import (
	"context"

	"github.com/vladesco/e-commerce/internal/ddd"
)

type DomainEventHandlers interface {
	OnOrderCreated(ctx context.Context, event ddd.Event) error
	OnStoreParticipationEnabled(ctx context.Context, event ddd.Event) error
	OnStoreParticipationDisabled(ctx context.Context, event ddd.Event) error
	OnProductAdded(ctx context.Context, event ddd.Event) error
	OnProductRemoved(ctx context.Context, event ddd.Event) error
}
