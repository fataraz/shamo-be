package http

import (
	"github.com/labstack/echo/v4"
	"net/http"

	ctxSess "shamo-be/internal/shared/utils/context"
)

type AppContext struct {
	echo.Context
	CtxSess *ctxSess.Context
}

func Parse(c echo.Context) *AppContext {
	data := c.Get(ctxSess.AppSession)
	ctxSess := data.(*ctxSess.Context)
	return &AppContext{Context: c, CtxSess: ctxSess}
}

func (c *AppContext) Ok(resp interface{}) error {
	var data interface{}
	data = resp
	if data == nil {
		data = struct{}{}
	}
	c.CtxSess.SetResponseCode(http.StatusOK)
	c.CtxSess.Lv4(data)

	return c.Context.JSON(http.StatusOK, resp)
}
