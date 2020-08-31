package conflux

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/models"
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
	addr, err := models.GenerateConfluxAddress()
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}
	cfxTransaction.Address = string(addr)
	//insert into database here
	err = cfxTransaction.Create(db)
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}
	fmt.Println(cfxTransaction)

	var data = map[string]interface{}{
		"address": addr,
	}

	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"success": true,
	})
}
