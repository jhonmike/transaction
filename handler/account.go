package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonmike/transaction/commons"
	"github.com/jhonmike/transaction/model"
	"github.com/jinzhu/gorm"
)

func createAccountHandler(accountResource model.AccountResource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var account model.Account
		if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
			commons.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		account, err := accountResource.CreateAccount(account)
		if err != nil {
			commons.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		commons.RespondJSON(w, http.StatusCreated, account)
	}
}

func getAccountHandler(accountResource model.AccountResource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		accountID := vars["accountID"]
		account, err := accountResource.GetAccountByID(model.Account{}, accountID)
		if err != nil {
			commons.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		commons.RespondJSON(w, http.StatusOK, account)
	}
}

// MakeAccountHandlers Adds the account module handlers to their endpoints
func MakeAccountHandlers(r *mux.Router, db *gorm.DB) {
	accountResource := model.NewAccountResource(db)

	r.HandleFunc("/accounts", createAccountHandler(accountResource)).Methods("POST")
	r.HandleFunc("/accounts/{accountID}", getAccountHandler(accountResource)).Methods("GET")
}
