package stellar

import (
	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/config"
)

type Env struct {
	Config *config.Config
	Router fiber.Router
}

func (s *Env) SetupRoutes() {
	stellar := s.Router.Group("/federation")
	stellar.Get("/", s.get)
	stellar.Post("/", s.createAddr)
}
