package domain

import (
	"github.com/google/uuid"
	"github.com/stackus/errors"
)

var (
	ErrProductNameIsMissed    = errors.Wrap(errors.ErrBadRequest, "[PRODUCT]: product name is missed")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "[PRODUCT]: product price can`t be negative")
)

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
