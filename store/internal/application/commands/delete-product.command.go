package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type DeleteProductCommand struct {
	ProductId string
}

type DeleteProductHandler struct {
	productRepository domain.ProductRepository
}

func NewDeleteProductHandler(productRepository domain.ProductRepository) *DeleteProductHandler {
	return &DeleteProductHandler{
		productRepository,
	}
}

func (handler *DeleteProductHandler) DeleteProduct(ctx context.Context, command DeleteProductCommand) error {
	return errors.Wrap(handler.productRepository.Delete(ctx, command.ProductId), "error deleting product")
}
