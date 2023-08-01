package container

import (
	"shamo-be/internal/infrastructure/postgres"
	database "shamo-be/internal/shared/database"
	productSvc "shamo-be/internal/usecase/product"
)

// Container ...
type Container struct {
	ProducSvc productSvc.Service
}

// Setup ...
func Setup() *Container {
	db := database.New()

	// repository
	productRepo := postgres.NewProductsRepository(db)

	// service
	productSvc := productSvc.New(productRepo)
	return &Container{
		ProducSvc: productSvc,
	}
}
