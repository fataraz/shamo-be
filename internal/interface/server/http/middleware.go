package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shamo-be/internal/interface/interceptor"
	"shamo-be/internal/shared/config"
	"shamo-be/internal/shared/logger"
	"shamo-be/internal/shared/utils"
	"shamo-be/internal/shared/utils/context"
)

// setupMiddleware ...
func setupMiddleware(server *echo.Echo, cfg *config.Config) {
	server.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId := c.Request().Header.Get(echo.HeaderXRequestID)
			if len(reqId) == 0 {
				reqId = utils.GenerateThreadId()
			}

			ctxSess := context.New(logger.GetLogger()).
				SetXRequestID(reqId).
				SetAppName(cfg.Apps.Name).
				SetAppVersion(cfg.Apps.Version).
				SetPort(cfg.Apps.HttpPort).
				SetSrcIP(c.RealIP()).
				SetURL(c.Request().URL.String()).
				SetMethod(c.Request().Method).
				SetHeader(c.Request().Header)

			ctxSess.Lv1("Incoming Request")

			c.Set(context.AppSession, ctxSess)

			return h(c)
		}
	})

	interceptor := interceptor.New()
	server.Use(interceptor.ValidateAccess())

	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.PATCH, echo.DELETE, echo.HEAD, echo.OPTIONS},
		AllowHeaders: []string{
			"Content-Type", "Origin", "Accept", "Authorization", "Content-Length", "X-Requested-With",
			"OS-Type", "Device-Name", "Device-Serial", "OS-Version", "App-Version",
		},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}
