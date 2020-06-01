package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/jhonmike/transaction/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTransactionResource struct {
	mock.Mock
}

func (m *mockTransactionResource) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	args := m.Called(transaction)
	return args.Get(0).(model.Transaction), args.Error(1)
}

type mockOperationTypeResource struct {
	mock.Mock
}

func (m *mockOperationTypeResource) GetOperationTypeByID(operationType model.OperationType, ID interface{}) (model.OperationType, error) {
	args := m.Called(operationType)
	return args.Get(0).(model.OperationType), args.Error(1)
}

func TestCreateTransactionHandler(t *testing.T) {
	spyTransaction := model.Transaction{
		ID:              1,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		AccountID:       1,
		OperationTypeID: 1,
		Description:     "description...",
		Amount:          123.45,
	}

	rjson, _ := json.Marshal(model.Transaction{
		AccountID:       spyTransaction.AccountID,
		OperationTypeID: spyTransaction.OperationTypeID,
		Description:     spyTransaction.Description,
		Amount:          spyTransaction.Amount,
	})
	req, err := http.NewRequest("POST", "/transactions", strings.NewReader(string(rjson)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	transactionResource := new(mockTransactionResource)
	transactionResource.On("CreateTransaction", mock.Anything).
		Return(spyTransaction, nil)

	operationTypeResource := new(mockOperationTypeResource)
	operationTypeResource.On("GetOperationTypeByID", mock.Anything).
		Return(model.OperationType{
			ID:          spyTransaction.OperationTypeID,
			Description: model.CashPurchase,
		}, nil)

	handler := http.HandlerFunc(createTransactionHandler(transactionResource, operationTypeResource))

	handler.ServeHTTP(rr, req)

	var transaction model.Transaction
	json.NewDecoder(rr.Body).Decode(&transaction)
	assert.Equal(t, spyTransaction.AccountID, transaction.AccountID, "return new transaction with account id value")
	assert.Equal(t, spyTransaction.OperationTypeID, transaction.OperationTypeID, "return new transaction with operation type id value")
	assert.Equal(t, spyTransaction.Description, transaction.Description, "return new transaction with description value")
	assert.Equal(t, spyTransaction.Amount, transaction.Amount, "return new transaction with amount value")
	assert.Equal(t, http.StatusCreated, rr.Code, "should return status 201")
}
