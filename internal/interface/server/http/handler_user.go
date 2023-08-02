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
	context := Parse(c)
	ctxSess := context.CtxSess
	req := &user.RegisterReq{}
	if err := c.Bind(req); err != nil {
		ctxSess.ErrorMessage = err.Error()
		ctxSess.Lv4()

		errors := FormatValidationError(err)
		errorMessage := map[string]any{"error": errors}
		resp := APIResponse("register account failed", http.StatusBadRequest, "error", errorMessage)
		return c.JSON(http.StatusBadRequest, resp)
	}

	ctxSess.Request = req
	// create user account
	if err := u.service.RegisterUser(ctxSess, req); err != nil {
		ctxSess.Lv4()

		errorMessage := map[string]any{"error": err.Error()}
		resp := APIResponse("register account failed", http.StatusBadRequest, "error", errorMessage)
		return c.JSON(http.StatusBadRequest, resp)
	}
	resp := APIResponse("account has been registered", http.StatusOK, "success", nil)
	return c.JSON(http.StatusOK, resp)
}
