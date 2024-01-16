package domain

import "context"

type ProductRepository interface {
	Save(ctx context.Context, store *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, productId string) error
	Find(ctx context.Context, productId string) (*Product, error)
	FindAll(ctx context.Context, storeId string) ([]*Product, error)
}
