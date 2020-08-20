package models

import (
	"database/sql"
	"time"
)

//Transaction model containing details of transaction
type Transaction struct {
	ID          int64
	AcnNo       string
	Bank        string
	Sender      *string
	SenderEmail *string
	Amount      int
	Currency    string
	Coin        int
	Crypto      string
	createdAt   time.Time
}

func (t *Transaction) Create(db *sql.DB) (int64, int64, error) {
	sqlQuery := "INSERT transactions SET accountNumber = ?, bank = ?, sender = ?, senderEmail = ?, amount = ?, currency = ?, coin = ?, crypto = ?, createdAt = ?"

	stmt, err := db.Prepare(sqlQuery)
	defer stmt.Close()
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(t.AcnNo, t.Bank, t.Sender, t.SenderEmail, t.Amount, t.Currency, t.Coin, t.Crypto)
	if err != nil {
		return 0, 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	return rowsAffected, lastInsertedID, err
}

func (t *Transaction) Get(db *sql.DB) ([]Transaction, error) {
	sqlQuery := "SELECT * FROM transactions"
	stmt, err := db.Prepare(sqlQuery)
	var result []Transaction
	defer stmt.Close()

	if err != nil {
		return result, err
	}
	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		return result, err
	}

	for rows.Next() {
		tran := new(Transaction)
		err := rows.Scan(&tran)
		if err != nil {
			return nil, err
		}
		result = append(result, *tran)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, err
}
