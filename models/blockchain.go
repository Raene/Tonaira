package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func BlockchainAddress(xpub string, apiKey string) (map[string]interface{}, error) {
	var result map[string]interface{}

	baseUrl, err := url.Parse("https://api.blockchain.info/v2/receive")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return result, errors.New("Blockchain Url is bad")
	}

	params := url.Values{}

	params.Add("xpub", xpub)
	params.Add("callback", "https://mystore.com?invoice_id=058921123")
	params.Add("key", apiKey)
	params.Add("gap_limit", "100")
	baseUrl.RawQuery = params.Encode()

	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err)
		return result, errors.New("Blockchain Get APi error")
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return result, errors.New("Blockchain Response body error")
	}

	json.Unmarshal([]byte(data), &result)

	return result, nil
}

func MonitorBlockChainAddress(address string, apiKey string) error {

	str := fmt.Sprintf(`{"key":"%s","addr":"%s","callback":"http://localhost:3000/api/v1/blockchain/init-transfer","onNotification":"DELETE", "op":"RECEIVE"}`, apiKey, address)

	resp, err := http.Post("https://api.blockchain.info/v2/receive/balance_update", "text/plain", bytes.NewBufferString(str))
	if err != nil {
		// handle err
		return err
	}
	defer resp.Body.Close()
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// fmt.Println(string(body))

	return nil
}
