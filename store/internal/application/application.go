package application

import (
	"github.com/vladesco/e-commerce/store/internal/application/commands"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type Application struct {
	applicationCommands
}

type applicationCommands struct {
	*commands.CreateProductHandler
	*commands.CreateStoreHandler
}

func NewApplication(storeRepository domain.StoreRepository, productRepository domain.ProductRepository) *Application {
	return &Application{
		applicationCommands{
			commands.NewCreateProductHandler(storeRepository, productRepository),
			commands.NewCreateStoreHandler(storeRepository),
		},
	}
}
