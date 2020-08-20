package transaction

import (
	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/models"
)

func (t *Transaction) get(ctx *fiber.Ctx) {
	db := t.Config.Db
	transactions := models.Transaction{}

	trans, err := transactions.Get(db)
	if err != nil {
		ctx.Next(err)
		return
	}

	var data = map[string]interface{}{
		"transactions": trans,
	}
	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"message": "success",
	})
}
