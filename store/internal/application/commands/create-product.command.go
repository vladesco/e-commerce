package commands

import (
	"context"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/internal/ddd"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type CreateProductCommand struct {
	StoreId     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

type CreateProductHandler struct {
	storeRepository   domain.StoreRepository
	productRepository domain.ProductRepository
	domainPublisher   ddd.EventPublisher
}

func NewCreateProductHandler(storeRepository domain.StoreRepository, productRepository domain.ProductRepository, domainPublisher ddd.EventPublisher) *CreateProductHandler {
	return &CreateProductHandler{
		storeRepository,
		productRepository,
		domainPublisher,
	}
}

func (handler *CreateProductHandler) CreateProduct(ctx context.Context, command CreateProductCommand) error {
	_, err := handler.storeRepository.Find(ctx, command.StoreId)

	if err != nil {
		return errors.Wrap(err, "error finding store")
	}

	product, err := domain.CreateProduct(command.StoreId, command.Name, command.Description, command.SKU, command.Price)

	if err != nil {
		return errors.Wrap(err, "error creating product")
	}

	err = handler.productRepository.Save(ctx, product)

	if err != nil {
		return errors.Wrap(err, "error saving product")
	}

	return errors.Wrap(handler.domainPublisher.Publish(ctx, product.GetEvents()...), "error publishing product events")
}
