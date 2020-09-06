package models

import (
	"fmt"
	"net/http"

	"github.com/rpip/paystack-go"
)

/*MakeTransfer initiates a transfer to the specified User in the Transaction Type, and amount is the amount in Naira
 */
func MakeTransfer(t Transaction, amount float32) (*paystack.Transfer, error) {
	apiKey := "sk_test_c6cb616397d40d3972fb01d1c7c2a7f9875ab5af"
	client := paystack.NewClient(apiKey, http.DefaultClient)

	banks, err := client.Bank.List()
	if err != nil || !(len(banks.Values) > 0) {
		fmt.Printf("Expected Bank list, got %d, returned error %v", len(banks.Values), err)
	}

	var bankCode string
	//for comparism, convert both strings bank name and t.Bank to lowerCase leters
	for _, bank := range banks.Values {
		if bank.Slug == t.Bank {
			bankCode = bank.Code
			break
		}
	}

	receipient := &paystack.TransferRecipient{
		Type:          "Nuban",
		AccountNumber: t.AccountNumber,
		BankCode:      bankCode,
		Currency:      "NGN",
	}

	receipient1, err := client.Transfer.CreateRecipient(receipient)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	req := &paystack.TransferRequest{
		Source:    "balance",
		Reason:    "tonaira.com",
		Amount:    amount,
		Recipient: receipient1.RecipientCode,
	}

	transfer, err := client.Transfer.Initiate(req)
	if err != nil {
		return nil, err
	}

	//store transfer details in received db
	//tx_hash, crypto amount received, cryptoCurrency
	//fmt.Println(transfer.Status)
	return transfer, nil
}
