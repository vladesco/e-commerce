package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/internal/ddd"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type EnableParticipationCommand struct {
	StoreId string
}

type EnableParticipationHandler struct {
	storeRepository domain.StoreRepository
	domainPublisher ddd.EventPublisher
}

func NewEnableParticipationHandler(storeRepository domain.StoreRepository, domainPublisher ddd.EventPublisher) *EnableParticipationHandler {
	return &EnableParticipationHandler{
		storeRepository,
		domainPublisher,
	}
}

func (handler *EnableParticipationHandler) EnableParticipation(ctx context.Context, command EnableParticipationCommand) error {
	store, err := handler.storeRepository.Find(ctx, command.StoreId)

	if err != nil {
		return errors.Wrap(err, "error finding store")
	}

	if err = store.EnableParticipation(); err != nil {
		return errors.Wrap(err, "error enabling participation")

	}

	if err = handler.storeRepository.Update(ctx, store); err != nil {
		return errors.Wrap(err, "error updating store")
	}

	return errors.Wrap(handler.domainPublisher.Publish(ctx, store.GetEvents()...), "error publishing store events")
}
