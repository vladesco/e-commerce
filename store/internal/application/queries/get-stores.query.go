package queries

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type GetStoresQuery struct {
}

type GetStoresHandler struct {
	storeRepository domain.StoreRepository
}

func NewGetStoresHandler(storeRepository domain.StoreRepository) *GetStoresHandler {
	return &GetStoresHandler{
		storeRepository,
	}
}

func (handler *GetStoresHandler) GetStores(ctx context.Context, query GetStoresQuery) ([]*domain.Store, error) {
	storeList, err := handler.storeRepository.FindAll(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error getting stores")
	}

	return storeList, nil
}
