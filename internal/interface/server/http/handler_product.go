package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shamo-be/internal/usecase/product"
)

// productHandler ...
type productHandler struct {
	service product.Service
}

// SetupProductHandler ...
func SetupProductHandler(service product.Service) *productHandler {
	handler := &productHandler{
		service,
	}
	if handler.service == nil {
		panic("service is nil")
	}
	return handler
}

// getProducts : get all items of products
func (p *productHandler) getProducts(c echo.Context) error {
	products, err := p.service.FindProducts()
	if err != nil {
		errors := FormatValidationError(err)
		response := APIResponse("failed to get products", http.StatusBadRequest, "error", errors)
		c.JSON(http.StatusBadRequest, response)
	}
	response := APIResponse("products", http.StatusOK, "success", products)
	return c.JSON(http.StatusOK, response)
}
