package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Accounts struct {
	gorm.Model
	DocumentNumber string
	Transactions   []Transactions `gorm:"foreignkey:AccountID"`
}

type OperationsTypes struct {
	gorm.Model
	Description  string
	Transactions []Transactions `gorm:"foreignkey:OperationsTypeID"`
}

type Transactions struct {
	gorm.Model
	AccountID        uint
	OperationsTypeID uint
	Description      string
	Amount           uint
	EventDate        time.Time
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
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=transaction password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&Accounts{}, &OperationsTypes{}, &Transactions{})

	r := mux.NewRouter()
	r.HandleFunc("/accounts", createAccountHandler).Methods("POST")
	r.HandleFunc("/accounts/{accountID}", getAccountHandler).Methods("GET")
	r.HandleFunc("/transactions", createTransactionHandler).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
