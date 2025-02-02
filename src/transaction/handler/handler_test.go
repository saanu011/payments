package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"payments/src/transaction/domain"
)

func TestHandler_GetTransactionDetailsByID(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/transaction/%s", "101"), nil).WithContext(context.Background())
	r = mux.SetURLVars(r, map[string]string{
		"transactionID": "101",
	})
	data := &domain.Transaction{
		ID:              "101",
		AccountID:       "11",
		OperationTypeID: 1,
		Amount:          100,
	}

	expectedBody := `{"ID":"101","AccountID":"11","OperationTypeID":1,"Amount":100,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"}`

	w := httptest.NewRecorder()

	txnService := &MockService{}
	handler := NewHandler(txnService)
	txnService.On("FindTransactionByID", r.Context(), "101").Return(data, nil).Once()

	handler.GetTransactionDetailsByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expectedBody, w.Body.String())

}

func TestHandler_CreateTransaction(t *testing.T) {
	data := &domain.Transaction{
		ID:              "101",
		AccountID:       "11",
		OperationTypeID: 1,
		Amount:          100,
	}

	reqBody, _ := json.Marshal(data)

	r := httptest.NewRequest(http.MethodPost, "/transaction",
		bytes.NewReader(reqBody)).WithContext(context.Background())


	expectedBody := `{"ID":"101","AccountID":"11","OperationTypeID":1,"Amount":100,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z"}`

	w := httptest.NewRecorder()

	txnService := &MockService{}
	handler := NewHandler(txnService)
	txnService.On("CreateTransaction", r.Context(), data).Return(data, nil).Once()

	handler.CreateTransaction(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expectedBody, w.Body.String())

}
