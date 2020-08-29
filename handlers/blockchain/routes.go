package blockchain

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/config"
	"github.com/spf13/viper"
)

type Env struct {
	Config *config.Config
	Router fiber.Router
	ApiKey string
	Xpub   string
}

func (e *Env) SetupRoutes() {
	viper.SetConfigFile(".env")
	// Find and read the config file
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	e.ApiKey = viper.Get("BLOCKCHAIN_APIKEY").(string)
	e.Xpub = viper.Get("BLOCKCHAIN_XPUB").(string)

	conflux := e.Router.Group("/blockchain")
	conflux.Get("/", e.getAddr)
}
