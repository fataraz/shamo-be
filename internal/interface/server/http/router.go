package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func SetupRouter(server *echo.Echo, handler *Handler) {
	server.GET("/ping", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "services up and running... "+time.Now().Format(time.RFC3339))
	})

	api := server.Group("/api")

	// product
	product := api.Group("/products")
	product.GET("", handler.productHandler.getProducts)
}
