package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/application"
	"github.com/vladesco/e-commerce/store/internal/application/commands"
)

type createStoreDTO struct {
	Name string `json:"name"`
}

func createStoreHandler(app *application.Application, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var dto createStoreDTO

	json.Unmarshal(body, &dto)

	err = app.CreateStoreHandler.CreateStore(r.Context(), commands.CreateStoreCommand(dto))

	if err != nil {
		http.Error(w, errors.Unwrap(err).Error(), errors.HTTPCode(err))
	}
}
