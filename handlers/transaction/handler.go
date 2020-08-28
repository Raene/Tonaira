package transaction

import (
	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/models"
)

func (t *Transaction) get(ctx *fiber.Ctx) {
	db := t.Config.Db
	transactions := models.Transaction{}

	trans, errs := transactions.Get(db)
	if len(errs) > 0 {
		ctx.Status(500).JSON(fiber.Map{
			"data":    errs,
			"message": "failed",
		})
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
