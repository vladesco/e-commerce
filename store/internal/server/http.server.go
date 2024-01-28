package server

import (
	"fmt"
	"net/http"

	"github.com/vladesco/e-commerce/internal/logger"
	"github.com/vladesco/e-commerce/store/internal/application"
)

type ServerConfig struct {
	Port int
}

func RegisterServer(logger *logger.Logger, app *application.Application, config *ServerConfig) error {
	mux := http.NewServeMux()

	registerStoreHandlers(mux, app)
	registerProductHandlers(mux, app)

	return http.ListenAndServe(fmt.Sprintf(":%d", config.Port), mux)
}

func registerStoreHandlers(mux *http.ServeMux, app *application.Application) {
	mux.HandleFunc("/store", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createStoreHandler(app, w, r)
		case http.MethodDelete:
			disableParticipationHandler(app, w, r)
		case http.MethodPatch:
			enableParticipationHandler(app, w, r)
		default:
			http.Error(w, "Unsupported Method", http.StatusBadRequest)
		}
	})
}

func registerProductHandlers(mux *http.ServeMux, app *application.Application) {
	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createProductHandler(app, w, r)
		case http.MethodDelete:
			deleteProductHandler(app, w, r)
		default:
			http.Error(w, "Unsupported Method", http.StatusBadRequest)
		}
	})
}
