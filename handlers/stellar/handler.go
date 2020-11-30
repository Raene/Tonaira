package stellar

import (
	"strings"

	"github.com/gofiber/fiber"
	"github.com/raene/Tonaira/models"
)

func (s *Env) get(ctx *fiber.Ctx) {
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

func (s *Env) createAddr(ctx *fiber.Ctx) {
	db := s.Config.Db
	stellarTransaction := models.Transaction{}
	stellarUser := models.StellarUser{}

	err := ctx.BodyParser(&stellarTransaction)
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	v := s.Config.Val
	if ok, errors := models.ValidateInputs(stellarTransaction, v); !ok {
		ctx.Status(500).JSON(&fiber.Map{
			"data":    errors,
			"success": false,
		})
		return
	}

	stellarUser.AccountId = stellarTransaction.Sender
	stellarTransaction.Address, err = models.CreateStellarUser(stellarUser, db)
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	err = stellarTransaction.Create(db)
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	var data = map[string]interface{}{
		"address": stellarTransaction.Address,
	}

	ctx.Status(200).JSON(fiber.Map{
		"data":    data,
		"success": true,
	})
}
