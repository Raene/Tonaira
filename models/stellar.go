package models

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/protocols/horizon/operations"
)

type StellarUser struct {
	ID             int64  `json:"id"`
	AccountId      string `json:accountId`
	StellarAddress string `json:stellarAddress`
	MemoType       string `json:memoType`
	Memo           string `json:memo`
}

//StellarStream opens a stream to the specified account and listens for payment
func StellarStream() {

	client := horizonclient.DefaultTestNetClient
	opRequest := horizonclient.OperationRequest{ForAccount: "GBNOVJRIUHB3S4BIMVJPE3LJSYNHOOZURIQLFDO2N2JY2OPOXG5KWJP2"}
	ctx := context.Background()
	printHandler := func(op operations.Operation) {
		fmt.Println("Watching now")
		fmt.Println(op)
	}
	fmt.Println("Watching now")
	err := client.StreamPayments(ctx, opRequest, printHandler)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateStellarUser(user StellarUser, db *gorm.DB) (string, error) {
	user.AccountId = strings.ToLower(user.AccountId)

	re := regexp.MustCompile(`\s+`)
	user.AccountId = re.ReplaceAllString(user.AccountId, "_")

	err := db.Create(user).Error
	if err != nil {
		return "", err
	}
	url := "tonaira.com"
	return user.AccountId + "*" + url, nil
}

func GetRecordByName(db *gorm.DB, name string) (StellarUser, []error) {
	user := StellarUser{}
	errs := db.Where("stellar_address = ?", name).First(&user).GetErrors()
	if len(errs) != 0 {
		return user, errs
	}
	return user, nil
}

func GetRecordByAccountId(db *gorm.DB, accountId string) (StellarUser, []error) {
	user := StellarUser{}
	errs := db.Where("account_id = ?", accountId).First(&user).GetErrors()
	if len(errs) != 0 {
		return user, errs
	}
	return user, nil
}
