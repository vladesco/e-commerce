package queries

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type GetCatalogQuery struct {
	StoreId string
}

type GetCatalogQueryHandler struct {
	productRepository domain.ProductRepository
}

func NewGetCatalogQueryHandler(productRepository domain.ProductRepository) *GetCatalogQueryHandler {
	return &GetCatalogQueryHandler{
		productRepository,
	}
}

func (handler *GetCatalogQueryHandler) GetCatalog(ctx context.Context, query GetCatalogQuery) ([]*domain.Product, error) {
	products, err := handler.productRepository.FindAll(ctx, query.StoreId)

	if err != nil {
		return nil, errors.Wrap(err, "error getting catalog")
	}

	return products, nil
}
