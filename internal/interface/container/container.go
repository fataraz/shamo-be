package container

import (
	"shamo-be/internal/infrastructure/postgres"
	"shamo-be/internal/shared/config"
	database "shamo-be/internal/shared/database"
	authSvc "shamo-be/internal/usecase/auth"
	productSvc "shamo-be/internal/usecase/product"
	userSvc "shamo-be/internal/usecase/user"
)

// Container ...
type Container struct {
	Config    *config.Config
	ProducSvc productSvc.Service
	UserSvc   userSvc.Service
	AuthSvc   authSvc.Service
}

// Setup ...
func Setup() *Container {
	// Construct Config
	cfg := config.NewConfig("./resources/config.json")

	db, err := database.New(cfg.Database.Db)
	if err != nil {
		panic(err)
	}

	// repository
	productRepo := postgres.NewProductsRepository(db)
	userRepo := postgres.NewUsersRepository(db)

	// service
	productSvc := productSvc.New(productRepo)
	userSvc := userSvc.New(userRepo)
	authSvc := authSvc.New(userRepo)

	return &Container{
		Config:    cfg,
		ProducSvc: productSvc,
		UserSvc:   userSvc,
		AuthSvc:   authSvc,
	}
}
