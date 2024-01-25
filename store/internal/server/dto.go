package server

type CreateProductDTO struct {
	StoreId     string  `json:"storeId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Price       float64 `json:"price"`
}

type CreateStoreDTO struct {
	Name string `json:"name"`
}
