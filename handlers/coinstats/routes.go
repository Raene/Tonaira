//Package coinstats is a package that deals with fetching stats for specified coins
package coinstats

import (
	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/config"
)

//CoinStats struct used to define values needed for this package
type CoinStats struct {
	Config *config.Config
	Router fiber.Router
}

//SetupRoutes assembles all the routes for this handler
func (c *CoinStats) SetupRoutes() {
	coinstats := c.Router.Group("/coin-stats")
	coinstats.Get("/", c.getStats)
}
