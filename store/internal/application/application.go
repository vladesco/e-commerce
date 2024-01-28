package application

import (
	"github.com/vladesco/e-commerce/store/internal/application/commands"
	"github.com/vladesco/e-commerce/store/internal/application/queries"
	"github.com/vladesco/e-commerce/store/internal/domain"
)

type Application struct {
	applicationCommands
	applicationQueries
}

type applicationCommands struct {
	*commands.CreateProductHandler
	*commands.DeleteProductHandler
	*commands.CreateStoreHandler
	*commands.EnableParticipationHandler
	*commands.DisableParticipationHandler
}

type applicationQueries struct {
	*queries.GetCatalogHandler
	*queries.GetProductHandler
	*queries.GetStoreHandler
	*queries.GetStoresHandler
}

func NewApplication(storeRepository domain.StoreRepository, productRepository domain.ProductRepository) *Application {
	return &Application{
		applicationCommands{
			commands.NewCreateProductHandler(storeRepository, productRepository),
			commands.NewDeleteProductHandler(productRepository),
			commands.NewCreateStoreHandler(storeRepository),
			commands.NewEnableParticipationHandler(storeRepository),
			commands.NewDisableParticipationHandler(storeRepository),
		},
		applicationQueries{
			queries.NewGetCatalogHandler(productRepository),
			queries.NewGetProductHandler(productRepository),
			queries.NewGetStoreHandler(storeRepository),
			queries.NewGetStoresHandler(storeRepository),
		},
	}
}
