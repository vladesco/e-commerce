package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/internal/ddd"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type CreateStoreCommand struct {
	Name string
}

type CreateStoreHandler struct {
	storeRepository domain.StoreRepository
	domainPublisher ddd.EventPublisher
}

func NewCreateStoreHandler(storeRepository domain.StoreRepository, domainPublisher ddd.EventPublisher) *CreateStoreHandler {
	return &CreateStoreHandler{
		storeRepository,
		domainPublisher,
	}
}

func (handler *CreateStoreHandler) CreateStore(ctx context.Context, command CreateStoreCommand) error {
	store, err := domain.CreateStore(command.Name)

	if err != nil {
		return errors.Wrap(err, "error creating store")
	}

	if err = handler.storeRepository.Save(ctx, store); err != nil {
		return errors.Wrap(err, "error saving store")
	}

	return errors.Wrap(handler.domainPublisher.Publish(ctx, store.GetEvents()...), "error publushin store event")
}
