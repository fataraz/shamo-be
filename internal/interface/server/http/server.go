package http

import (
	"github.com/labstack/echo/v4"

	"shamo-be/internal/interface/container"
)

// StartHttpService ...
func StartHttpService(cont *container.Container) {
	server := echo.New()
	server.HideBanner = true

	setupMiddleware(server, cont.Config)
	SetupRouter(server, SetupHandler(cont))

	// Start server http
	server.Logger.Fatal(server.Start(cont.Config.AppAdress()))
}
