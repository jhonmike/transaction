package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Transaction model that represents each transaction executed in the account
type Transaction struct {
	ID              uint       `gorm:"primary_key" json:"id"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `sql:"index" json:"deleted_at"`
	AccountID       uint       `json:"account_id"`
	OperationTypeID uint       `json:"operation_type_id"`
	Description     string     `json:"description"`
	Amount          float64    `json:"amount"`
	EventDate       time.Time  `json:"event_date"`
}

// TransactionResource ...
type TransactionResource interface {
	CreateTransaction(transaction Transaction) (Transaction, error)
}

type transactionResource struct {
	db *gorm.DB
}

// NewTransactionResource ...
func NewTransactionResource(db *gorm.DB) TransactionResource {
	return transactionResource{db}
}

func (resource transactionResource) CreateTransaction(transaction Transaction) (Transaction, error) {
	if err := resource.db.Create(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
