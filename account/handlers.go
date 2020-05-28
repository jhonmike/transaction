package account

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// Account model representing the customer is back account
type Account struct {
	gorm.Model
	DocumentNumber string
	Transactions   []Transaction `gorm:"foreignkey:AccountID"`
}

// OperationType model that represents the types of operations transacted in the account
type OperationType struct {
	gorm.Model
	Description  string
	Transactions []Transaction `gorm:"foreignkey:OperationTypeID"`
}

// Transaction model that represents each transaction executed in the account
type Transaction struct {
	gorm.Model
	AccountID       uint
	OperationTypeID uint
	Description     string
	Amount          uint
	EventDate       time.Time
}

func createAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "document_number": "12345678900" }`)
}

func getAccountHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "account_id": "%v", "document_number": "12345678900" }`, vars["accountID"])
}

func createTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "account_id": "1", "operation_type_id": "4", "amount": "123.45" }`)
}

// MakeAccountHandlers Adds the account module handlers to their endpoints
func MakeAccountHandlers(r *mux.Router) {
	r.HandleFunc("/accounts", createAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{accountID}", getAccountHandler).Methods("GET")
	r.HandleFunc("/transactions", createTransactionHandler).Methods("POST")
}
