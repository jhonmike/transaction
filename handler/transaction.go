package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jhonmike/transaction/commons"
	"github.com/jhonmike/transaction/model"
	"github.com/jinzhu/gorm"
)

// TransactionValidate valid with error messages the struct
type TransactionValidate struct {
	Error map[string]string `json:"error"`
}

func (t *TransactionValidate) isValid(transaction model.Transaction) bool {
	t.Error = make(map[string]string)
	isValid := true

	if transaction.Amount <= 0 {
		isValid = false
		t.Error["amount"] = "To create a transaction you need the amount to be greater than zero"
	}

	if transaction.AccountID == 0 {
		isValid = false
		t.Error["account_id"] = "To create a transaction you need the accountID"
	}

	if transaction.OperationTypeID == 0 {
		isValid = false
		t.Error["operation_type_id"] = "To create a transaction you need the Operation Type ID"
	}

	return isValid
}

func createTransactionHandler(transactionResource model.TransactionResource, operationTypeResource model.OperationTypeResource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction model.Transaction
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			commons.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}

		validate := TransactionValidate{}
		if !validate.isValid(transaction) {
			commons.RespondJSON(w, http.StatusBadRequest, validate)
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
