package transactions

import (
	Db "Challenge/internal/adapters/db"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

var db, _ = Db.NewDatabaseConnection()

func (s *DefaultTransactionService) GetAllTransactions() ([]Transaction, error) {
	var trans = []Transaction{}
	db.Find(&trans)
	if trans != nil {

		return trans, nil
	} else {

		return nil, errors.New("no transactions found")
	}
}

func (s *DefaultTransactionService) CreateTransaction(trans *Transaction) (Transaction, error) {
	trans.Id = uuid.New()
	trans.Createdat = time.Now().String()
	trans.Status = false
	if err := db.Create(&trans).Error; err != nil {

		return *trans, err
	}

	//Produce Kafa Event after create transaction
	Produce(trans)

	return *trans, nil
}

func (s *DefaultTransactionService) UpdateTransaction(JsonTransaction string) (bool, error) {
	var transaction Transaction
	json.Unmarshal([]byte(JsonTransaction), &transaction)
	transaction.Status = true
	if err := db.Model(&transaction).Update("status", true).Error; err != nil {
		return true, err
	}
	return false, nil
}
