package http

import "shamo-be/internal/interface/container"

// Handler ...
type Handler struct {
	productHandler *productHandler
	userHandler    *userHandler
	authHandler    *authHandler
}

// SetupHandler ...
func SetupHandler(container *container.Container) *Handler {
	productHandler := SetupProductHandler(container.ProducSvc)
	userHandler := SetupUserHandler(container.UserSvc)
	authHandler := SetupAuthHandler(container.AuthSvc)
	return &Handler{
		productHandler: productHandler,
		userHandler:    userHandler,
		authHandler:    authHandler,
	}
}
