package commands

import (
	"context"

	"github.com/stackus/errors"
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
}

func NewCreateProductHandler(storeRepository domain.StoreRepository, productRepository domain.ProductRepository) *CreateProductHandler {
	return &CreateProductHandler{
		storeRepository,
		productRepository,
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

	return errors.Wrap(handler.productRepository.Save(ctx, product), "error saving product")
}
