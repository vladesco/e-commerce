package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/application"
	"github.com/vladesco/e-commerce/store/internal/application/commands"
)

type deleteProductDTO struct {
	ProductId string `json:"productId"`
}

func deleteProductHandler(app *application.Application, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var dto deleteProductDTO

	json.Unmarshal(body, &dto)

	err = app.DeleteProductHandler.DeleteProduct(r.Context(), commands.DeleteProductCommand(dto))

	if err != nil {
		http.Error(w, errors.Unwrap(err).Error(), errors.HTTPCode(err))
	}
}
