package repository

import (
	"context"

	"github.com/vladesco/e-commerce/internal/logger"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type StoreRepository struct {
	logger *logger.Logger
}

func NewStoreRepository(logger *logger.Logger) *StoreRepository {
	return &StoreRepository{
		logger,
	}
}

func (repository *StoreRepository) Save(ctx context.Context, store *domain.Store) error {
	repository.logger.Debug().Msgf("saving store %+v", store)
	return nil
}

func (repository *StoreRepository) Update(ctx context.Context, store *domain.Store) error {
	repository.logger.Debug().Msgf("updating store %+v", store)
	return nil
}

func (repository *StoreRepository) Delete(ctx context.Context, storeId string) error {
	repository.logger.Debug().Msgf("deleteing store %s", storeId)
	return nil
}

func (repository *StoreRepository) Find(ctx context.Context, storeId string) (*domain.Store, error) {
	repository.logger.Debug().Msgf("finding store %s", storeId)
	return nil, nil
}

func (repository *StoreRepository) FindAll(ctx context.Context) ([]*domain.Store, error) {
	repository.logger.Debug().Msg("finding all stores")
	return nil, nil
}
