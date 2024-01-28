package queries

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type GetProductQuery struct {
	ProductId string
}

type GetProductQueryHandler struct {
	productRepository domain.ProductRepository
}

func NewGetProductQueryHandler(productRepository domain.ProductRepository) *GetProductQueryHandler {
	return &GetProductQueryHandler{
		productRepository,
	}
}

func (handler *GetProductQueryHandler) GetProduct(ctx context.Context, query GetProductQuery) (*domain.Product, error) {
	product, err := handler.productRepository.Find(ctx, query.ProductId)

	if err != nil {
		return nil, errors.Wrap(err, "error getting product")
	}

	return product, nil
}
