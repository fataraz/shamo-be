package http

import "shamo-be/internal/interface/container"

// Handler ...
type Handler struct {
	productHandler *productHandler
}

// SetupHandler ...
func SetupHandler(container *container.Container) *Handler {
	productHandler := SetupProductHandler(container.ProducSvc)
	return &Handler{
		productHandler: productHandler,
	}
}
