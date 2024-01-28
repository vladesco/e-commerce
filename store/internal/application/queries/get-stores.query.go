package queries

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type GetStoresQuery struct {
}

type GetStoresQueryHandler struct {
	storeRepository domain.StoreRepository
}

func NewGetStoresQueryHandler(storeRepository domain.StoreRepository) *GetStoresQueryHandler {
	return &GetStoresQueryHandler{
		storeRepository,
	}
}

func (handler *GetStoresQueryHandler) GetStores(ctx context.Context, query GetStoresQuery) ([]*domain.Store, error) {
	storeList, err := handler.storeRepository.FindAll(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error getting stores")
	}

	return storeList, nil
}
