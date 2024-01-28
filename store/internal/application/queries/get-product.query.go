package queries

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type GetProductQuery struct {
	ProductId string
}

type GetProductHandler struct {
	productRepository domain.ProductRepository
}

func NewGetProductHandler(productRepository domain.ProductRepository) *GetProductHandler {
	return &GetProductHandler{
		productRepository,
	}
}

func (handler *GetProductHandler) GetProduct(ctx context.Context, query GetProductQuery) (*domain.Product, error) {
	product, err := handler.productRepository.Find(ctx, query.ProductId)

	if err != nil {
		return nil, errors.Wrap(err, "error getting product")
	}

	return product, nil
}
