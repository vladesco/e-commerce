package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/stackus/errors"
	"github.com/vladesco/e-commerce/store/internal/application"
	"github.com/vladesco/e-commerce/store/internal/application/commands"
)

type enableParticipationDTO struct {
	StoreId string `json:"storeId"`
}

func enableParticipationHandler(app *application.Application, w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var dto enableParticipationDTO

	json.Unmarshal(body, &dto)

	err = app.EnableParticipationHandler.EnableParticipation(r.Context(), commands.EnableParticipationCommand(dto))

	if err != nil {
		http.Error(w, errors.Unwrap(err).Error(), errors.HTTPCode(err))
	}
}
