package main

import (
	"fmt"

	"github.com/raene/Tonaira/models"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	Config "github.com/raene/Tonaira/config"
	"github.com/raene/Tonaira/database"
	"github.com/raene/Tonaira/handlers/blockchain"
	"github.com/raene/Tonaira/handlers/coinstats"
	"github.com/raene/Tonaira/handlers/conflux"
	"github.com/raene/Tonaira/handlers/paystack"
)

//Routes interface every route should implement to get spawned
type Routes interface {
	SetupRoutes()
}

func spawnRoutes(m chan string, r ...Routes) {
	for _, v := range r {
		v.SetupRoutes()
	}
	m <- "Routes Setup"
}

func main() {
	var m chan string = make(chan string)
	db := database.Init()

	app := fiber.New()
	api := app.Group("/api/v1", logger.New())
	config := Config.Init(db)

	coinRoutes := &coinstats.CoinStats{Config: config, Router: api}

	confluxRoutes := &conflux.Env{
		Config: config,
		Router: api,
	}

	paystackRoutes := &paystack.Env{
		Config: config,
		Router: api,
	}

	blockchainRoutes := &blockchain.Env{
		Config: config,
		Router: api,
	}

	go spawnRoutes(m, coinRoutes, blockchainRoutes, confluxRoutes, paystackRoutes)
	go models.SpawnConfluxCron(db)

	fmt.Println(<-m)
	app.Listen(3000)

}
