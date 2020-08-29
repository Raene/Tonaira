package blockchain

import (
	"github.com/raene/Tonaira/models"

	//"strings"
	"github.com/gofiber/fiber"
)

func (e *Env) getAddr(ctx *fiber.Ctx) {
	db := e.Config.Db
	cfxTransaction := models.Transaction{}

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

	err = cfxTransaction.Create(db)
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	var data = map[string]interface{}{
		"address": cfxTransaction.Address,
	}

	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"success": true,
	})
}
