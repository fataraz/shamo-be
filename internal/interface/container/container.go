package container

import (
	"shamo-be/internal/infrastructure/postgres"
	database "shamo-be/internal/shared/database"
	authSvc "shamo-be/internal/usecase/auth"
	productSvc "shamo-be/internal/usecase/product"
	userSvc "shamo-be/internal/usecase/user"
)

// Container ...
type Container struct {
	ProducSvc productSvc.Service
	UserSvc   userSvc.Service
	AuthSvc   authSvc.Service
}

// Setup ...
func Setup() *Container {
	db := database.New()

	// repository
	productRepo := postgres.NewProductsRepository(db)
	userRepo := postgres.NewUsersRepository(db)

	// service
	productSvc := productSvc.New(productRepo)
	userSvc := userSvc.New(userRepo)
	authSvc := authSvc.New(userRepo)

	return &Container{
		ProducSvc: productSvc,
		UserSvc:   userSvc,
		AuthSvc:   authSvc,
	}
}
