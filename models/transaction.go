package models

import (
	"time"

	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/jinzhu/gorm"
)

//Transaction model containing details of transaction
type Transaction struct {
	ID            int64
	AccountNumber string
	Bank          string
	Sender        *string
	SenderEmail   *string
	Amount        int
	Network       string
	Address       types.Address
	createdAt     time.Time
}

func (t *Transaction) Create(db *gorm.DB) error {
	err := db.Create(&t).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) Get(db *gorm.DB) ([]Transaction, []error) {
	transactions := []Transaction{}
	errs := db.Find(&transactions).GetErrors()
	if len(errs) != 0 {
		return transactions, errs
	}
	return transactions, nil
}
