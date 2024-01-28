package queries

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type GetStoreQuery struct {
	StoreId string
}

type GetStoreQueryHandler struct {
	storeRepository domain.StoreRepository
}

func NewGetStoreQueryHandler(storeRepository domain.StoreRepository) *GetStoreQueryHandler {
	return &GetStoreQueryHandler{
		storeRepository,
	}
}

func (handler *GetStoreQueryHandler) GetStore(ctx context.Context, query GetStoreQuery) (*domain.Store, error) {
	store, err := handler.storeRepository.Find(ctx, query.StoreId)

	if err != nil {
		return nil, errors.Wrap(err, "error getting store")
	}

	return store, nil
}
