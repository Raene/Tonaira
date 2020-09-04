package paystack

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
	Viper  *viper.Viper
}

func (e *Env) SetupRoutes() {
	e.Viper = viper.New()
	e.Viper.SetConfigFile(".env")
	err := e.Viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	e.ApiKey = viper.Get("PAYSTACK_SECRET").(string)

	payStack := e.Router.Group("/paystack")
	payStack.Post("/verify-account", e.verifyAccount)
}
