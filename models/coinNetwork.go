package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Network struct {
	Key string
	Url  string
}

type Result struct {
    Payload map[string]string
    Error error
}

//Networks is a map of accepted Crypto networks and their various URLs
var Networks = map[string]Network{
	"btc": {Url: fmt.Sprintf("https://api.coinbase.com/v2/prices/BTC-USD/buy"), Key: "amount"},
	"xlm": {Url: fmt.Sprintf("https://api.coinbase.com/v2/prices/XLM-USD/buy"),Key: "amount"},
	"ngn":{Url: fmt.Sprintf("https://api.currencyfreaks.com/latest?apikey=80b78db254e64fd9a652b0a94b2331c1&symbols=NGN"),Key:"NGN"},
	//"cfx": {Url: fmt.Sprintf("https://www.worldcoinindex.com/apiservice/ticker?key=HSgipENwBaSTTnytJxMjtJf1lqU4QSNVHtM&label=cfxbtc&fiat=USD"),Key: "cfx"},
}

func CoinExchangeRate()  <-chan Result{
	x := make(chan Result)
	go func(){
		defer close(x)
		results,err := Results(Networks)
		result := Result{results,err}
		if result.Error != nil {
			x <- result
		}
		x <- result
	}()
	return x
}
//https://api.currencyfreaks.com/latest?apikey=80b78db254e64fd9a652b0a94b2331c1&symbols=NGN
func RatesQuery(url string) (map[string]interface{},error){
	var result map[string]interface{}
	response, err := http.Get(url)
	 if err != nil {
		 fmt.Println(err)
		 return nil,err
	}
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal([]byte(data), &result)
	return result,nil
}

func Results(networks map[string]Network) (map[string]string,error) {
	results := make(map[string]string)
	for k := range networks {
		url := networks[k].Url
		result,err := RatesQuery(url)
		if err != nil {
			fmt.Println(err)
			return results,err
	   }
		results[k] = Loop(result,networks[k].Key)
	}
	return results,nil
}

func Loop(result map[string]interface{}, k string) string {
	for key := range result {
		_,ok := result[k]
		if !ok {
			fmt.Println(k)
			fmt.Println(result)
			if _,ok = result[key].(map[string]interface{});!ok{
				continue
			}
			return Loop(result[key].(map[string]interface{}),k)
		}
 	  return result[k].(string)
	} 
	return ""
}
