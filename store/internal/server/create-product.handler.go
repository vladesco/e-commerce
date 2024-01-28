package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/application"
	"github.com/vladesco/e-commerce/store/internal/application/commands"
)

type createProductDTO struct {
	StoreId     string  `json:"storeId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Price       float64 `json:"price"`
}

func createProductHandler(app *application.Application, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var dto createProductDTO

	json.Unmarshal(body, &dto)

	err = app.CreateProductHandler.CreateProduct(r.Context(), commands.CreateProductCommand(dto))

	if err != nil {
		http.Error(w, errors.Unwrap(err).Error(), errors.HTTPCode(err))
	}
}
