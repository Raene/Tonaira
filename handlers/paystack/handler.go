package paystack

import (
	"fmt"

	"net/http"

	"github.com/gofiber/fiber"
	"github.com/rpip/paystack-go"
)

func (e *Env) verifyAccount(ctx *fiber.Ctx) {
	type account struct {
		Number   string `json:number`
		BankName string `json:bankName`
	}

	accn := account{}

	err := ctx.BodyParser(&accn)
	if err != nil {
		ctx.Status(400).JSON(&fiber.Map{
			"data":    err,
			"success": false,
		})
		return
	}

	client := paystack.NewClient(e.ApiKey, http.DefaultClient)

	banks, err := client.Bank.List()
	if err != nil || !(len(banks.Values) > 0) {
		fmt.Printf("Expected Bank list, got %d, returned error %v", len(banks.Values), err)
	}
	var bankCode string
	//for comparism, convert both strings bank name and t.Bank to lowerCase leters
	for _, bank := range banks.Values {
		if bank.Slug == accn.BankName {
			bankCode = bank.Code
			break
		}
	}

	resp, err := client.Bank.ResolveAccountNumber(accn.Number, bankCode)

	if err != nil {
		fmt.Printf("Expected error, got %+v'", err)
	}

	ctx.Status(200).JSON(fiber.Map{
		"data":    resp,
		"success": true,
	})
}
