package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Transaction model containing details of transaction
type Transaction struct {
	ID            int64   `json:"id,omitempty" validate:"omitempty"`
	AccountNumber string  `json:"accountNumber" validate:"required"`
	Bank          string  `json:"bank" validate:"required"`
	Sender        string  `json:"sender,omitempty" validate:"omitempty"`
	SenderEmail   *string `json:"senderEmail" validate:"required,email"`
	ExchangeRate  float32 `json:"exchangeRate" validate:"required"`
	Network       string  `json:"network" validate:"required"`
	Naira         float32 `json:"naira" validate:"required"`
	Status        bool		//transaction status is to be updated when we have paid the client
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

func (t *Transaction) GetBy(db *gorm.DB) (*Transaction, error) {
	transaction := Transaction{}
	err := db.Where(&t).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
