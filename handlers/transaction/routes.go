package transaction

import (
	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/config"
)

type Transaction struct {
	Config *config.Config
	Router fiber.Router
}

func (t *Transaction) SetupRoutes() {
	transaction := t.Router.Group("/transaction")
	transaction.Get("/", t.get)
}
