package coinstats

import (
	"fmt"

	Coin "github.com/raene/Tonaira/models"

	"github.com/gofiber/fiber"
)

/*
this handler receives the amount to be converted in USD and returns the converted value in btc
along with the current exchange rate
*/
func (c *CoinStats) getStats(ctx *fiber.Ctx) {
	// var e chan error = make(chan error)
	var x chan map[string]map[string]interface{} = make(chan map[string]map[string]interface{})
	// var err error

	go Coin.CoinExchangeRate(x)

	// err = <-e
	// if err != nil {
	// 	ctx.Next(err)
	// 	return
	// }

	xChangeRate := <-x
	fmt.Println(xChangeRate)

	var data = map[string]interface{}{
		"ExchangeRate": xChangeRate,
	}
	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"message": "success",
	})
}
