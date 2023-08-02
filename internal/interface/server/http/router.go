package http

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// SetupRouter ...
func SetupRouter(server *echo.Echo, handler *Handler) {
	server.GET("/ping", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "services up and running... "+time.Now().Format(time.RFC3339))
	})
	server.GET("/", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "OK")
	})

	api := server.Group("/api")

	// register
	api.POST("/register", handler.userHandler.registerUser)

	// auth
	api.POST("/login", handler.authHandler.login)

	// product
	product := api.Group("/products")
	product.GET("", handler.productHandler.getProducts)
}
