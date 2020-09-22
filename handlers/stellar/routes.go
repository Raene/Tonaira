package stellar

import (
	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/config"
)

type Stellar struct {
	Config *config.Config
	Router fiber.Router
}

func (s *Stellar) SetupRoutes() {
	stellar := s.Router.Group("/federation")
	stellar.Get("/", s.get)
}
