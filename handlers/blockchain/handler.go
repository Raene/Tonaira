package blockchain

import (
	"bytes"	
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	//"strings"
	"github.com/gofiber/fiber"
)

func (e *Env) getAddr(ctx *fiber.Ctx) {
	// db := e.Config.Db
	// cfxTransaction := models.Transaction{}

	// err := ctx.BodyParser(&cfxTransaction)
	// if err != nil {
	// 	ctx.Status(400).JSON(&fiber.Map{
	// 		"data":    err,
	// 		"success": false,
	// 	})
	// 	return
	// }

	baseUrl, err := url.Parse("https://api.blockchain.info/v2/receive")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return
	}

	params := url.Values{}

	params.Add("xpub", e.Xpub)
	params.Add("callback", "https://mystore.com?invoice_id=058921123")
	params.Add("key", e.ApiKey)
	params.Add("gap_limit", "100")
	baseUrl.RawQuery = params.Encode()
	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	var result map[string]interface{}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal([]byte(data), &result)

fmt.Println(result)
// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
//chamge callback url
str := fmt.Sprintf(`{"key":"%s","addr":"%s","callback":"https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7","onNotification":"DELETE", "op":"RECEIVE"}`,e.ApiKey,result["address"].(string))

resp, err := http.Post("https://api.blockchain.info/v2/receive/balance_update","text/plain",bytes.NewBufferString(str))
if err != nil {
	// handle err
fmt.Println(err)
}
defer resp.Body.Close()	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
fmt.Println(string(body))
}
