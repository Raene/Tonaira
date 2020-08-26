package conflux

import (
	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/config"
)

type Env struct {
	Config *config.Config
	Router fiber.Router
}

func (e *Env) SetupRoutes() {
	conflux := e.Router.Group("/conflux")
	conflux.Post("/", e.getAddr)
}
