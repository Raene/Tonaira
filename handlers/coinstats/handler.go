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
	// var r chan Coin.Result = make(chan Coin.Result)

	r := Coin.CoinExchangeRate()
	result := <-r
	if result.Error != nil {
		fmt.Println(result.Error)
		ctx.Next(result.Error)
		return
	}
	
	xChangeRate := result.Payload
	fmt.Println(xChangeRate)
	fmt.Println(result.Payload)

	var data = map[string]interface{}{
		"ExchangeRate": xChangeRate,
	}
	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"message": "success",
	})
}
