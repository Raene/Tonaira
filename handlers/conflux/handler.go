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
			"success": false,
			"message": err,
		})
		return
	}
	addr, err := models.GenerateConfluxAddress()
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	cfxTransaction.Address = addr
	//insert into database here
	err = cfxTransaction.Create(db)
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
		return
	}
	fmt.Println(cfxTransaction)

	var data = map[string]interface{}{
		"address": addr,
	}

	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"message": "success",
	})
}
