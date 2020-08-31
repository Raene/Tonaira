package blockchain

import (
	"fmt"

	"github.com/raene/Tonaira/models"
	Coin "github.com/raene/Tonaira/models"

	//"strings"
	"github.com/gofiber/fiber"
)

func (e *Env) getAddr(ctx *fiber.Ctx) {
	db := e.Config.Db
	cfxTransaction := models.Transaction{}

	var x chan interface{} = make(chan interface{})
	go Coin.CoinExchangeRate(x)

	err := ctx.BodyParser(&cfxTransaction)
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	result, err := models.BlockchainAddress(e.Xpub, e.ApiKey)
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	cfxTransaction.Address = result["address"].(string)

	err = models.MonitorBlockChainAddress(cfxTransaction.Address, e.ApiKey)
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	rate := <-x
	cfxTransaction.ExchangeRate = rate.(float32)
	err = cfxTransaction.Create(db)
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	var data = map[string]interface{}{
		"address":      cfxTransaction.Address,
		"exchangeRate": cfxTransaction.ExchangeRate,
	}

	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"success": true,
	})
}

func (e *Env) initTransfer(ctx *fiber.Ctx) {
	type CallbackPayload struct {
		Address string `json:address`
		Value   int    `json:value`
		TxHash  string `json:transactionHash`
		Status  bool   `json:status`
	}

	db := e.Config.Db
	var payload CallbackPayload

	err := ctx.BodyParser(&payload)
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	//initial payload value is in satoshi, so convert to btc
	payload.Value = payload.Value / 100000000

	t := models.Transaction{}

	err = db.Where("address =?", payload.Address).First(&t).Error
	if err != nil {
		fmt.Println(err)
		return
	}

	amount := t.ExchangeRate * float32(payload.Value)
	//convert amount to naira first
	ngn := amount * 40000
	_, err = models.MakeTransfer(t, ngn)
	if err != nil {
		fmt.Println(err)
		payload.Status = false
		err = db.Create(&payload).Error
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// if transfer.Status != " " {
	// 	fmt.Println("in here")
	// 	payload.Status = false
	// 	err = db.Create(&payload).Error
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }

	// payload.Status = true
	// err = db.Create(&payload).Error
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}
