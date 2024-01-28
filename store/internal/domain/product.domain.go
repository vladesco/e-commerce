package domain

import (
	"context"

	"github.com/google/uuid"
	"github.com/stackus/errors"
)

type ProductRepository interface {
	Save(ctx context.Context, store *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, productId string) error
	Find(ctx context.Context, productId string) (*Product, error)
	FindAll(ctx context.Context, storeId string) ([]*Product, error)
}

type Product struct {
	Id          string
	StoreId     string
	Name        string
	Description string
	SKU         string
	Price       float64
}

func CreateProduct(storeId, name, description, sku string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameIsMissed
	}

	if price < 0 {
		return nil, ErrProductPriceIsNegative
	}

	product := &Product{
		uuid.NewString(),
		storeId,
		name,
		description,
		sku,
		price,
	}

	return product, nil
}

var (
	ErrProductNameIsMissed    = errors.Wrap(errors.ErrBadRequest, "[PRODUCT]: product name is missed")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "[PRODUCT]: product price can`t be negative")
)
