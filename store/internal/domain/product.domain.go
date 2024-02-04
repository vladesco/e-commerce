package domain

import (
	"github.com/google/uuid"
	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/internal/ddd"
)

type Product struct {
	ddd.AggregateBase
	id          string
	storeId     string
	name        string
	description string
	sku         string
	price       float64
}

func CreateProduct(storeId, name, description, sku string, price float64) (*Product, error) {
	if name == "" {
		return nil, ErrProductNameIsMissed
	}

	if price < 0 {
		return nil, ErrProductPriceIsNegative
	}

	product := &Product{
		ddd.CrateAggregate(),
		uuid.NewString(),
		storeId,
		name,
		description,
		sku,
		price,
	}

	product.AddEvent(&ProductCreated{product})

	return product, nil
}

func (product *Product) Remove() error {
	product.AddEvent(&ProductDeleted{product})

	return nil
}

var (
	ErrProductNameIsMissed    = errors.Wrap(errors.ErrBadRequest, "[PRODUCT]: product name is missed")
	ErrProductPriceIsNegative = errors.Wrap(errors.ErrBadRequest, "[PRODUCT]: product price can`t be negative")
)
