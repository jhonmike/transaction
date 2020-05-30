package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jhonmike/transaction/model"
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/assert"
)

func newDB() *gorm.DB {
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	db, _ := gorm.Open(mocket.DriverName, "connection_string")
	return db
}

func TestCreateAccountHandler(t *testing.T) {
	documentNumber := "666"

	payload := fmt.Sprintf(`{ "document_number": "%s" }`, documentNumber)
	req, err := http.NewRequest("POST", "/accounts", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	db := newDB()

	handler := http.HandlerFunc(createAccountHandler(db))
	handler.ServeHTTP(rr, req)

	var account model.Account
	json.NewDecoder(rr.Body).Decode(&account)
	assert.Equal(t, documentNumber, account.DocumentNumber, "return new account with Document Number")
	assert.Equal(t, http.StatusCreated, rr.Code, "should return status 201")
}

func TestGetAccountHandler(t *testing.T) {
	accountID := 666
	documentNumber := "999"

	req, err := http.NewRequest("GET", fmt.Sprintf("/accounts/%s", string(accountID)), nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	db := newDB()
	query := `SELECT * FROM "accounts"  WHERE "accounts"."deleted_at" IS NULL AND (("accounts"."document_number" = ʚ)) ORDER BY "accounts"."id" ASC LIMIT 1`
	commonReply := []map[string]interface{}{{"id": accountID, "document_number": documentNumber}}
	mocket.Catcher.Reset().NewMock().
		WithQuery(query).
		WithReply(commonReply)

	router := mux.NewRouter()
	router.HandleFunc("/accounts/{accountID}", getAccountHandler(db))
	router.ServeHTTP(rr, req)

	var account model.Account
	json.NewDecoder(rr.Body).Decode(&account)
	assert.Equal(t, documentNumber, account.DocumentNumber, "return new account with account ID")
	assert.Equal(t, http.StatusOK, rr.Code, "should return status 200")
}

func TestCreateTransactionHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/transactions", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createTransactionHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{ "account_id": "1", "operation_type_id": "4", "amount": "123.45" }`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
