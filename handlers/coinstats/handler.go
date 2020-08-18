package coinstats

import (
	"encoding/json"
	"errors"
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

func (c *coinQuery) coinToCurrency(s chan string, e chan error) {
	var network = map[string]Network{
		"btc": {name: "bitcoin", url: fmt.Sprintf("https://blockchain.info/tobtc?currency=%s&value=%f", c.currency, c.amount)},
	}

	// we can use this if statement to check to see if
	// a given key exists within a map in Go
	netwk, ok := network[c.name]
	if !ok {
		e <- errors.New("Unregistered Crypto Network")
		s <- " "
		return
	}

	response, err := http.Get(netwk.url)
	if err != nil {
		e <- err
		s <- " "
		return
	}

	data, _ := ioutil.ReadAll(response.Body)
	e <- nil
	s <- string(data)

}

func coinExchangeRate(x chan interface{}) {
	type exchangeRates struct {
		fifteenMins float64 `json:"15m"`
		last        float64
		buy         float64
		sell        float64
		symbol      string
	}

	var result map[string]interface{}

	response, _ := http.Get("https://blockchain.info/ticker")
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(data), &result)
	x <- result["USD"].(map[string]interface{})["buy"]

}

type coinQuery struct {
	currency string
	amount   float64
	name     string
}

func (c *CoinStats) getStats(ctx *fiber.Ctx) {
	var e chan error = make(chan error)
	var s chan string = make(chan string)
	var x chan interface{} = make(chan interface{})
	coin := new(coinQuery)
	var err error

	//clean query strings of empty space
	coin.currency = strings.TrimSpace(ctx.Query("currency"))
	coin.name = strings.TrimSpace(ctx.Query("name"))

	//convert amount query to Float64, just to make sure an amount is passed not a string
	coin.amount, err = strconv.ParseFloat(strings.TrimSpace(ctx.Query("amount")), 64)
	if err != nil {
		ctx.Next(err)
		return
	}

	go coin.coinToCurrency(s, e)
	go coinExchangeRate(x)

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
