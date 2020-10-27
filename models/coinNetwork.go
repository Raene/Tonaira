package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Network struct {
	Name string
	Url  string
}

//Networks is a map of accepted Crypto networks and their various URLs
var Networks = map[string]Network{
	"btc": {Name: "bitcoin", Url: fmt.Sprintf("https://api.coinbase.com/v2/prices/BTC-USD/buy")},
	"xlm": {Name: "stellar", Url: fmt.Sprintf("https://api.coinbase.com/v2/prices/XLM-USD/buy")},
}

func CoinExchangeRate(x chan map[string]map[string]interface{}) {

	results := make(map[string]map[string]interface{})
	for k, _ := range Networks {
		url := Networks[k].Url
		var result map[string]interface{}
		response, _ := http.Get(url)
		// if err != nil {
		// 	e <- err
		// 	return
		// }
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal([]byte(data), &result)
		results[k] = result
	}

	x <- results
	//e <- nil

}
