package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jhonmike/transaction/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

func main() {
	cfg := config.MustReadFromEnv()

	dbCfg := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbBase, cfg.DbPass)
	db, err := gorm.Open("postgres", dbCfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Account{}, &OperationType{}, &Transaction{})

	r := mux.NewRouter()
	r.HandleFunc("/accounts", createAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{accountID}", getAccountHandler).Methods("GET")
	r.HandleFunc("/transactions", createTransactionHandler).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", cfg.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
