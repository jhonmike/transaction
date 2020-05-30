package account

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonmike/transaction/commons"
	"github.com/jhonmike/transaction/model"
	"github.com/jinzhu/gorm"
)

func createAccountHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var account model.Account
		if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
			commons.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		db.Create(&account)

		commons.RespondJSON(w, http.StatusCreated, account)
	}
}

func getAccountHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var account model.Account
		if err := db.First(&account, model.Account{DocumentNumber: vars["accountID"]}).Error; err != nil {
			commons.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		commons.RespondJSON(w, http.StatusOK, account)
	}
}

func createTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{ "account_id": "1", "operation_type_id": "4", "amount": "123.45" }`)
}

// MakeAccountHandlers Adds the account module handlers to their endpoints
func MakeAccountHandlers(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/accounts", createAccountHandler(db)).Methods("POST")
	r.HandleFunc("/accounts/{accountID}", getAccountHandler(db)).Methods("GET")
	r.HandleFunc("/transactions", createTransactionHandler).Methods("POST")
}
