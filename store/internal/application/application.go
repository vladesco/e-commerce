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
	*commands.EnableParticipationCommandHandler
	*commands.DisableParticipationCommandHandler
}

type applicationQueries struct {
	*queries.GetCatalogQueryHandler
	*queries.GetProductQueryHandler
	*queries.GetStoreQueryHandler
	*queries.GetStoresQueryHandler
}

func NewApplication(storeRepository domain.StoreRepository, productRepository domain.ProductRepository) *Application {
	return &Application{
		applicationCommands{
			commands.NewCreateProductHandler(storeRepository, productRepository),
			commands.NewDeleteProductHandler(productRepository),
			commands.NewCreateStoreHandler(storeRepository),
			commands.NewEnableParticipationCommandHandler(storeRepository),
			commands.NewDisableParticipationCommandHandler(storeRepository),
		},
		applicationQueries{
			queries.NewGetCatalogQueryHandler(productRepository),
			queries.NewGetProductQueryHandler(productRepository),
			queries.NewGetStoreQueryHandler(storeRepository),
			queries.NewGetStoresQueryHandler(storeRepository),
		},
	}
}
