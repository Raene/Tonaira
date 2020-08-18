package coinstats

import (
	"strconv"
	"strings"

	Coin "github.com/raene/Tonaira/models"

	"github.com/gofiber/fiber"
)

/*
this handler receives the amount to be converted in USD and returns the converted value in btc
along with the current exchange rate
*/
func (c *CoinStats) getStats(ctx *fiber.Ctx) {
	var e chan error = make(chan error)
	var s chan string = make(chan string)
	var x chan interface{} = make(chan interface{})
	coin := new(Coin.CoinQuery)
	var err error

	//clean query strings of empty space
	coin.Currency = strings.TrimSpace(ctx.Query("currency"))
	coin.Name = strings.TrimSpace(ctx.Query("name"))

	//convert amount query to Float64, just to make sure an amount is passed not a string
	coin.Amount, err = strconv.ParseFloat(strings.TrimSpace(ctx.Query("amount")), 64)
	if err != nil {
		ctx.Next(err)
		return
	}

	go coin.CoinToCurrency(s, e)
	go Coin.CoinExchangeRate(x)

	err = <-e
	if err != nil {
		ctx.Next(err)
		return
	}

	coinValue := <-s
	xChangeRate := <-x

	var data = map[string]interface{}{
		"CoinValue":    coinValue,
		"ExchangeRate": xChangeRate,
	}
	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"message": "success",
	})
}
