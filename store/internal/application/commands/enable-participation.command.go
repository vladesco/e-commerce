package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type EnableParticipationCommand struct {
	StoreId string
}

type EnableParticipationHandler struct {
	storeRepository domain.StoreRepository
}

func NewEnableParticipationHandler(storeRepository domain.StoreRepository) *EnableParticipationHandler {
	return &EnableParticipationHandler{
		storeRepository,
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

	return errors.Wrap(handler.storeRepository.Update(ctx, store), "error updating store")
}
