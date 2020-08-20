package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	Config "github.com/raene/Tonaira/config"
	"github.com/raene/Tonaira/database"
	"github.com/raene/Tonaira/handlers/coinstats"
	"github.com/raene/Tonaira/handlers/transaction"
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
	transactionRoutes := &transaction.Transaction{
		Config: config,
		Router: api,
	}

	go spawnRoutes(m, coinRoutes, transactionRoutes)

	fmt.Println(<-m)
	app.Listen(3000)

}
