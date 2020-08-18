package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Network struct {
	name string
	url  string
}

type CoinQuery struct {
	Currency string
	Amount   float64
	Name     string
}

type exchangeRates struct {
	fifteenMins float64 `json:"15m"`
	last        float64
	buy         float64
	sell        float64
	symbol      string
}

func (c *CoinQuery) CoinToCurrency(s chan string, e chan error) {
	var network = map[string]Network{
		"btc": {name: "bitcoin", url: fmt.Sprintf("https://blockchain.info/tobtc?currency=%s&value=%f", c.Currency, c.Amount)},
	}

	// we can use this if statement to check to see if
	// a given key exists within a map in Go
	netwk, ok := network[c.Name]
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

func CoinExchangeRate(x chan interface{}) {
	var result map[string]interface{}

	response, _ := http.Get("https://blockchain.info/ticker")
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(data), &result)
	x <- result["USD"].(map[string]interface{})["buy"]

}
