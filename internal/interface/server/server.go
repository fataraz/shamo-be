package server

import (
	"shamo-be/internal/interface/container"
	"shamo-be/internal/interface/server/http"
)

// StartService ...
func StartService(container *container.Container) {
	// start http server
	http.StartHttpService(container)
}
