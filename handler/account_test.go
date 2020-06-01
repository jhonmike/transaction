package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/jhonmike/transaction/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAccountResource struct {
	mock.Mock
}

func (m *mockAccountResource) CreateAccount(account model.Account) (model.Account, error) {
	args := m.Called(account)
	return args.Get(0).(model.Account), args.Error(1)
}

func (m *mockAccountResource) GetAccountByID(account model.Account, ID interface{}) (model.Account, error) {
	args := m.Called(account)
	return args.Get(0).(model.Account), args.Error(1)
}

func TestCreateAccountHandler(t *testing.T) {
	spyAccount := model.Account{
		ID:             1,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		DocumentNumber: "666",
	}

	rjson, _ := json.Marshal(model.Account{
		DocumentNumber: spyAccount.DocumentNumber,
	})
	req, err := http.NewRequest("POST", "/accounts", strings.NewReader(string(rjson)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	accountResource := new(mockAccountResource)
	accountResource.On("CreateAccount", mock.Anything).Return(spyAccount, nil)

	handler := http.HandlerFunc(createAccountHandler(accountResource))
	handler.ServeHTTP(rr, req)

	var account model.Account
	json.NewDecoder(rr.Body).Decode(&account)
	assert.Equal(t, spyAccount.ID, account.ID, "return new account with ID")
	assert.Equal(t, spyAccount.DocumentNumber, account.DocumentNumber, "return new account with Document Number")
	assert.Equal(t, http.StatusCreated, rr.Code, "should return status 201")
	accountResource.AssertExpectations(t)
}

func TestCreateACcountHandlerWithoutDocumentNumber(t *testing.T) {
	rjson, _ := json.Marshal(model.Account{
		DocumentNumber: "",
	})
	req, err := http.NewRequest("POST", "/accounts", strings.NewReader(string(rjson)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	accountResource := new(mockAccountResource)
	handler := http.HandlerFunc(createAccountHandler(accountResource))
	handler.ServeHTTP(rr, req)

	var messageError AccountValidate
	json.NewDecoder(rr.Body).Decode(&messageError)
	assert.Equal(t, "To create a account you need your document number", messageError.Error["document_number"], "return error message")
	assert.Equal(t, http.StatusBadRequest, rr.Code, "should return status 400")
	accountResource.AssertExpectations(t)
}

func TestGetAccountHandler(t *testing.T) {
	spyAccount := model.Account{
		ID:             999,
		DocumentNumber: "666",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("/accounts/%d", spyAccount.ID), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	accountResource := new(mockAccountResource)
	accountResource.On("GetAccountByID", mock.Anything).Return(spyAccount, nil)

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{accountID}", getAccountHandler(accountResource))
	router.ServeHTTP(rr, req)

	var account model.Account
	json.NewDecoder(rr.Body).Decode(&account)
	assert.Equal(t, spyAccount.ID, account.ID, "return new account with account ID")
	assert.Equal(t, spyAccount.DocumentNumber, account.DocumentNumber, "return new account with Document Number")
	assert.Equal(t, http.StatusOK, rr.Code, "should return status 200")
	accountResource.AssertExpectations(t)
}
