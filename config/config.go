package config

import (
	"github.com/gofiber/fiber"
)

type Config struct {
	Router fiber.Router
}

func ConfigInit(router fiber.Router) *Config {
	config := Config{}
	config.Router = router

	return &config
}
