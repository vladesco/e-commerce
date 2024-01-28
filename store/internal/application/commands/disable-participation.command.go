package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type DisableParticipationCommand struct {
	StoreId string
}

type DisableParticipationHandler struct {
	storeRepository domain.StoreRepository
}

func NewDisableParticipationHandler(storeRepository domain.StoreRepository) *DisableParticipationHandler {
	return &DisableParticipationHandler{
		storeRepository,
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

	return errors.Wrap(handler.storeRepository.Update(ctx, store), "error updating store")
}
