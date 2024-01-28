package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type DisableParticipationCommand struct {
	StoreId string
}

type DisableParticipationCommandHandler struct {
	storeRepository domain.StoreRepository
}

func NewDisableParticipationCommandHandler(storeRepository domain.StoreRepository) *DisableParticipationCommandHandler {
	return &DisableParticipationCommandHandler{
		storeRepository,
	}
}

func (handler *DisableParticipationCommandHandler) DisableParticipation(ctx context.Context, command DisableParticipationCommand) error {
	store, err := handler.storeRepository.Find(ctx, command.StoreId)

	if err != nil {
		return errors.Wrap(err, "error finding store")
	}

	if err = store.DisableParticipation(); err != nil {
		return errors.Wrap(err, "error disabling participation")

	}

	return errors.Wrap(handler.storeRepository.Update(ctx, store), "error updating store")
}
