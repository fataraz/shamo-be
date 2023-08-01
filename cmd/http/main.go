package main

import (
	"shamo-be/internal/interface/container"
	"shamo-be/internal/interface/server"
)

// main ...
func main() {
	server.StartService(container.Setup())
}
