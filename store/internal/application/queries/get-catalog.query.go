package queries

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type GetCatalogQuery struct {
	StoreId string
}

type GetCatalogHandler struct {
	productRepository domain.ProductRepository
}

func NewGetCatalogHandler(productRepository domain.ProductRepository) *GetCatalogHandler {
	return &GetCatalogHandler{
		productRepository,
	}
}

func (handler *GetCatalogHandler) GetCatalog(ctx context.Context, query GetCatalogQuery) ([]*domain.Product, error) {
	productList, err := handler.productRepository.FindAll(ctx, query.StoreId)

	if err != nil {
		return nil, errors.Wrap(err, "error getting catalog")
	}

	return productList, nil
}
