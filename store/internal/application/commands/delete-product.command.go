package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/internal/ddd"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type DeleteProductCommand struct {
	ProductId string
}

type DeleteProductHandler struct {
	productRepository domain.ProductRepository
	domainPublisher   ddd.EventPublisher
}

func NewDeleteProductHandler(productRepository domain.ProductRepository, domainPublisher ddd.EventPublisher) *DeleteProductHandler {
	return &DeleteProductHandler{
		productRepository,
		domainPublisher,
	}
}

func (handler *DeleteProductHandler) DeleteProduct(ctx context.Context, command DeleteProductCommand) error {
	product, err := handler.productRepository.Find(ctx, command.ProductId)

	if err != nil {
		return errors.Wrap(err, "error finding product")
	}

	if err = product.Remove(); err != nil {
		return errors.Wrap(err, "error removing product")
	}

	if err = handler.productRepository.Delete(ctx, command.ProductId); err != nil {
		return errors.Wrap(err, "error deleting product")
	}

	return errors.Wrap(handler.domainPublisher.Publish(ctx, product.GetEvents()...), "error publishing product events")
}
