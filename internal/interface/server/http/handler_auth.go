package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shamo-be/internal/usecase/auth"
)

// authHandler...
type authHandler struct {
	service auth.Service
}

// SetupAuthHandler ...
func SetupAuthHandler(service auth.Service) *authHandler {
	handler := &authHandler{service: service}
	if handler.service == nil {
		panic("service is nil")
	}
	return handler
}

// login ...
func (a *authHandler) login(c echo.Context) error {
	context := Parse(c)
	ctxSess := context.CtxSess

	req := &auth.LoginReq{}
	if err := c.Bind(req); err != nil {
		ctxSess.ErrorMessage = err.Error()
		ctxSess.Lv4()
		errors := FormatValidationError(err)
		errorMessage := map[string]any{"error": errors}
		resp := APIResponse("login failed", http.StatusBadRequest, "error", errorMessage)
		return c.JSON(http.StatusBadRequest, resp)
	}

	login, err := a.service.Login(ctxSess, req)
	if err != nil {
		ctxSess.Lv4()
		errorMessage := map[string]any{"error": err.Error()}
		resp := APIResponse("login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := APIResponse("successfully logged in", http.StatusOK, "success", login)

	ctxSess.Lv4()
	return c.JSON(http.StatusOK, resp)
}
