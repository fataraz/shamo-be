package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

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
	context := Parse(c)
	ctxSess := context.CtxSess

	products, err := p.service.FindProducts(ctxSess)
	if err != nil {
		ctxSess.Lv4()
		resp := APIResponse("failed to get products", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := APIResponse("products", http.StatusOK, "success", products)
	ctxSess.Lv4(resp)

	return c.JSON(http.StatusOK, resp)
}
