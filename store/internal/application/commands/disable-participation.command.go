package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/internal/ddd"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type DisableParticipationCommand struct {
	StoreId string
}

type DisableParticipationHandler struct {
	storeRepository domain.StoreRepository
	domainPublisher ddd.EventPublisher
}

func NewDisableParticipationHandler(storeRepository domain.StoreRepository, domainPublisher ddd.EventPublisher) *DisableParticipationHandler {
	return &DisableParticipationHandler{
		storeRepository,
		domainPublisher,
	}
}

func (handler *DisableParticipationHandler) DisableParticipation(ctx context.Context, command DisableParticipationCommand) error {
	store, err := handler.storeRepository.Find(ctx, command.StoreId)

	if err != nil {
		return errors.Wrap(err, "error finding store")
	}

	if err = store.DisableParticipation(); err != nil {
		return errors.Wrap(err, "error disabling participation")

	}

	if err = handler.storeRepository.Update(ctx, store); err != nil {
		return errors.Wrap(err, "error updating store")
	}

	return errors.Wrap(handler.domainPublisher.Publish(ctx, store.GetEvents()...), "error publishing store events")
}
