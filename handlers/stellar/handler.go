package stellar

import (
	"strings"

	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/models"
)

func (s *Stellar) get(ctx *fiber.Ctx) {
	db := s.Config.Db
	q := strings.TrimSpace(ctx.Query("q"))
	t := strings.TrimSpace(ctx.Query("type"))

	if len(q) > 0 {
		ctx.Status(400).JSON(fiber.Map{
			"data":    "missing parameters, excepted parameter q",
			"success": false,
		})
	}

	if len(t) > 0 {
		ctx.Status(400).JSON(fiber.Map{
			"data":    "missing parameters, excepted parameter type",
			"success": false,
		})
	}

	var data models.StellarUser
	var err []error

	switch t {
	case "name":
		data, err = models.GetRecordByName(db, t)
	case "id":
		data, err = models.GetRecordByAccountId(db, t)
	}

	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"data":    err,
			"success": false,
		})
	}

	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"success": true,
	})
}

func (s *Stellar) createAddr(ctx *fiber.Ctx) {
	db := s.Config.Db
}
