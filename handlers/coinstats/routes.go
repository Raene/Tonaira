//Package coinstats is a package that deals with fetching stats for specified coins
package coinstats

import "github.com/raene/Tonaira/config"

//CoinStats struct used to define values needed for this package
type CoinStats struct {
	Config *config.Config
}

//SetupRoutes assembles all the routes for this handler
func (c *CoinStats) SetupRoutes() {
	coinstats := c.Config.Router.Group("/coin-stats")
	coinstats.Get("/", c.getStats)
}
