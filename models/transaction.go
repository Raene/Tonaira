package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Transaction model containing details of transaction
type Transaction struct {
	ID            int64   `json:"id"`
	AccountNumber string  `json:"accountNumber"`
	Bank          string  `json:"bank"`
	Sender        *string `json:"sender"`
	SenderEmail   *string `json:"senderEmail"`
	ExchangeRate  float32 `json:"exchangeRate"`
	Network       string  `json:"network"`
	Status        bool
	Address       string
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

func (t *Transaction) GetWhere(db *gorm.DB) ([]Transaction, []error) {
	transactions := []Transaction{}
	errs := db.Where("status = ? AND network =?", false, "cfx").Find(&transactions).GetErrors()
	if len(errs) != 0 {
		return transactions, errs
	}
	return transactions, nil

}

func (t *Transaction) Update(db *gorm.DB) error {
	err := db.Model(&t).Updates(&t).Error
	if err != nil {
		return err
	}
	return nil
}
