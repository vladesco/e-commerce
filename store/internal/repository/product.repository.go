package repository

import (
	"context"

	"github.com/vladesco/e-commerce/internal/logger"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type ProductRepository struct {
	logger *logger.Logger
}

func NewProductRepository(logger *logger.Logger) *ProductRepository {
	return &ProductRepository{
		logger,
	}
}

func (repository *ProductRepository) Save(ctx context.Context, product *domain.Product) error {
	repository.logger.Debug().Msgf("saving product %+v", product)
	return nil
}

func (repository *ProductRepository) Update(ctx context.Context, product *domain.Product) error {
	repository.logger.Debug().Msgf("updating product %+v", product)
	return nil
}

func (repository *ProductRepository) Delete(ctx context.Context, productId string) error {
	repository.logger.Debug().Msgf("deleteing product %s", productId)
	return nil
}

func (repository *ProductRepository) Find(ctx context.Context, productId string) (*domain.Product, error) {
	repository.logger.Debug().Msgf("finding product %s", productId)
	return nil, nil
}

func (repository *ProductRepository) FindAll(ctx context.Context, storeId string) ([]*domain.Product, error) {
	repository.logger.Debug().Msgf("finding all products %s", storeId)
	return nil, nil
}
