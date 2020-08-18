package coinstats

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber"
)

type Network struct {
	name string
	url  string
}

func coinToCurrency(curr string, amount float64, name string) (string, error) {
	var network = map[string]Network{
		"btc": {name: "bitcoin", url: fmt.Sprintf("https://blockchain.info/tobtc?currency=%s&value=%f", curr, amount)},
	}
	fmt.Println(name)
	// we can use this if statement to check to see if
	// a given key exists within a map in Go
	netwk, ok := network[name]
	if !ok {
		return "Unregistered Crypto Network", nil
	}

	response, err := http.Get(netwk.url)
	if err != nil {
		return "", err
	}

	data, _ := ioutil.ReadAll(response.Body)
	return string(data), nil

}

type coinQuery struct {
	currency string
	amount   float64
	name     string
}

func (c *CoinStats) getStats(ctx *fiber.Ctx) {
	coin := new(coinQuery)
	var err error
	//clean query strings of empty space
	coin.currency = strings.TrimSpace(ctx.Query("currency"))
	coin.name = strings.TrimSpace(ctx.Query("name"))

	//convert amount query to Float64, just to make sure an amount is passed not a string
	coin.amount, err = strconv.ParseFloat(strings.TrimSpace(ctx.Query("amount")), 64)
	if err != nil {
		ctx.Next(err)
	}

	data, err := coinToCurrency(coin.currency, coin.amount, coin.name)

	if err != nil {
		ctx.Next(err)
	}

	ctx.SendString(data)
	ctx.SendStatus(200)
}
