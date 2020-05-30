package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// Account model representing the customer is back account
type Account struct {
	gorm.Model
	DocumentNumber string        `json:"document_number"`
	Transactions   []Transaction `gorm:"foreignkey:AccountID" json:"transactions"`
}

// OperationType model that represents the types of operations transacted in the account
type OperationType struct {
	gorm.Model
	Description  string        `json:"description"`
	Transactions []Transaction `gorm:"foreignkey:OperationTypeID" json:"transactions"`
}

// Transaction model that represents each transaction executed in the account
type Transaction struct {
	gorm.Model
	AccountID       uint      `json:"account_id"`
	OperationTypeID uint      `json:"operation_type_id"`
	Description     string    `json:"description"`
	Amount          uint      `json:"amount"`
	EventDate       time.Time `json:"event_date"`
}

// NewDatabase created new instance to database...
func NewDatabase(host string, port string, user string, pass string, base string) *gorm.DB {
	cfg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, base, pass)
	db, err := gorm.Open("postgres", cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Account{}, &OperationType{}, &Transaction{})

	db.Create(&OperationType{Description: "Compra a Vista"})
	db.Create(&OperationType{Description: "Compra Parcelada"})
	db.Create(&OperationType{Description: "Saque"})
	db.Create(&OperationType{Description: "Pagamento"})

	return db
}
