package domain

type ProductCreated struct {
	Product *Product
}

func (ProductCreated) Name() string {
	return ProductCreatedEventName
}

type ProductDeleted struct {
	Product *Product
}

func (ProductDeleted) Name() string {
	return ProductDeletedEventName
}

var (
	ProductCreatedEventName = "product.ProductCreated"
	ProductDeletedEventName = "product.ProductDeleted"
)
