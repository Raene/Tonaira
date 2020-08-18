package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	Config "github.com/raene/Tonaira/config"
	"github.com/raene/Tonaira/handlers/coinstats"
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
	//var c chan *gorm.DB = make(chan *gorm.DB)
	var m chan string = make(chan string)
	//go database.InitDatabase(c)

	app := fiber.New()
	api := app.Group("/api/v1", logger.New())
	config := Config.ConfigInit(api)

	coinRoutes := &coinstats.CoinStats{Config: config}

	go spawnRoutes(m, coinRoutes)

	fmt.Println(<-m)
	app.Listen(3000)

}
