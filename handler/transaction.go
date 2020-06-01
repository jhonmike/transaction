package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonmike/transaction/commons"
	"github.com/jhonmike/transaction/model"
	"github.com/jinzhu/gorm"
)

func createTransactionHandler(transactionResource model.TransactionResource, operationTypeResource model.OperationTypeResource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction model.Transaction
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			commons.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		operationType, err := operationTypeResource.GetOperationTypeByID(model.OperationType{}, transaction.OperationTypeID)
		if err != nil {
			commons.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		if operationType.Description != model.Payment {
			transaction.Amount = -transaction.Amount
		}

		transaction, err = transactionResource.CreateTransaction(transaction)
		if err != nil {
			commons.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		commons.RespondJSON(w, http.StatusCreated, transaction)
	}
}

// MakeTransactionHandlers Adds the transaction handlers to their endpoints
func MakeTransactionHandlers(r *mux.Router, db *gorm.DB) {
	transactionResource := model.NewTransactionResource(db)
	operationTypeResource := model.NewOperationTypeResource(db)

	r.HandleFunc("/transactions", createTransactionHandler(transactionResource, operationTypeResource)).Methods("POST")
}
