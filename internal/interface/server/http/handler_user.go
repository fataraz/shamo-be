package http

import (
	"net/http"
	
	"github.com/labstack/echo/v4"

	"shamo-be/internal/usecase/user"
)

// userHandler ...
type userHandler struct {
	service user.Service
}

// SetupUserHandler ...
func SetupUserHandler(service user.Service) *userHandler {
	handler := &userHandler{service: service}
	if handler.service == nil {
		panic("service is nil")
	}
	return handler
}

// registerUser ...
func (u *userHandler) registerUser(c echo.Context) error {
	req := &user.RegisterReq{}
	if err := c.Bind(req); err != nil {
		errors := FormatValidationError(err)
		errorMessage := map[string]any{"error": errors}
		resp := APIResponse("register account failed", http.StatusBadRequest, "error", errorMessage)
		return c.JSON(http.StatusBadRequest, resp)
	}

	// create user account
	if err := u.service.RegisterUser(req); err != nil {
		resp := APIResponse("register account failed", http.StatusBadRequest, "error", nil)
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := APIResponse("account has been registered", http.StatusOK, "success", nil)
	return c.JSON(http.StatusOK, resp)
}
