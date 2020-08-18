package config

import (
	"github.com/gofiber/fiber"
)

//Config type contains Router
type Config struct {
	Router fiber.Router
}

//Init sets up the needed variables required for the server app to run
func Init(router fiber.Router) *Config {
	config := Config{}
	config.Router = router

	return &config
}
