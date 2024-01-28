package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type CreateStoreCommand struct {
	Name string
}

type CreateStoreHandler struct {
	storeRepository domain.StoreRepository
}

func NewCreateStoreHandler(storeRepository domain.StoreRepository) *CreateStoreHandler {
	return &CreateStoreHandler{
		storeRepository,
	}
}

func (handler *CreateStoreHandler) CreateStore(ctx context.Context, command CreateStoreCommand) error {
	store, err := domain.CreateStore(command.Name)

	if err != nil {
		return errors.Wrap(err, "error creating store")
	}

	return errors.Wrap(handler.storeRepository.Save(ctx, store), "error saving store")
}
