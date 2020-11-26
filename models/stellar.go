package models

import (
	"context"
	"fmt"
	"time"
	"regexp"
	"strings"
	"strconv"
	"math/rand"
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
func StellarStream(e chan error, db *gorm.DB) {

	client := horizonclient.DefaultTestNetClient
	//opRequest := horizonclient.OperationRequest{ForAccount: "GBNOVJRIUHB3S4BIMVJPE3LJSYNHOOZURIQLFDO2N2JY2OPOXG5KWJP2",Cursor:"1002793259241473"}
	opRequest := horizonclient.OperationRequest{ForAccount: "GBNOVJRIUHB3S4BIMVJPE3LJSYNHOOZURIQLFDO2N2JY2OPOXG5KWJP2"}
	ctx := context.Background()
	printHandler := func(op operations.Operation) {
		fmt.Println("Watching now")
		fmt.Println(op.GetType())
		fmt.Println(op.GetID())
		fmt.Println(op.GetTransactionHash())
		fmt.Println(op.IsTransactionSuccessful())
		fmt.Println(op.PagingToken())
		tx, err := client.TransactionDetail(op.GetTransactionHash())
		if err != nil {
			fmt.Println(err)
			StellarStream(e, db)
			e <- err
		}
		fmt.Println(tx.Memo)
		//fetch user from DB and start paystack
		record, errs := GetRecordByName(db, tx.Memo)
		if errs != nil {
			fmt.Println(errs)
			StellarStream(e, db)
			e <- errs[0]
		}
		t := Transaction{}
		t.Address = record.StellarAddress
		transaction, err := t.GetBy(db)
		if err != nil {
			fmt.Println(err)
			StellarStream(e, db)
			e <- err
		}
		fmt.Println(transaction)
	}
	fmt.Println("Watching now")
	err := client.StreamPayments(ctx, opRequest, printHandler)
	if err != nil {
		fmt.Println(err)
		StellarStream(e, db)
		e <- err
	}
	e <- nil
}

func CreateStellarUser(user StellarUser, db *gorm.DB) (string, error) {
	user.AccountId = strings.ToLower(user.AccountId)

	re := regexp.MustCompile(`\s+`)
	user.AccountId = re.ReplaceAllString(user.AccountId, "_")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	r1String := strconv.Itoa(r1.Intn(100))
	url := "tonaira.com"
	userUrl := user.AccountId + r1String + "*" + url

	user.StellarAddress = userUrl
	user.MemoType = "text"
	user.Memo = user.AccountId
	err := db.Create(&user).Error
	if err != nil {
		return "", err
	}
	return userUrl, nil
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
