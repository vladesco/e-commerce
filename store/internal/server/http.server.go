package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vladesco/e-commerce/internal/logger"
	"github.com/vladesco/e-commerce/store/internal/application"
	"github.com/vladesco/e-commerce/store/internal/application/commands"
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
			body, err := io.ReadAll(r.Body)

			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			var createStoreDTO CreateStoreDTO

			json.Unmarshal(body, &createStoreDTO)

			app.CreateStoreHandler.CreateStore(r.Context(), commands.CreateStoreCommand(createStoreDTO))
		default:
			http.Error(w, "Unsupported Method", http.StatusBadRequest)
		}
	})
}

func registerProductHandlers(mux *http.ServeMux, app *application.Application) {
	mux.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			body, err := io.ReadAll(r.Body)

			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			var createProductDTO CreateProductDTO

			json.Unmarshal(body, &createProductDTO)

			app.CreateProductHandler.CreateProduct(r.Context(), commands.CreateProductCommand(createProductDTO))
		default:
			http.Error(w, "Unsupported Method", http.StatusBadRequest)
		}
	})
}
