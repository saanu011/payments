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

	"payments/src/account/domain"
)

func TestHandler_GetAccountDetailsByID(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/account/%s", "101"), nil).WithContext(context.Background())
	r = mux.SetURLVars(r, map[string]string{
		"accountID": "101",
	})

	expectedBody := `{"id": "101","email":"test.mail","password":"******","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`

	w := httptest.NewRecorder()

	txnService := &MockService{}
	handler := NewHandler(txnService)
	txnService.On("FindAccountByID", r.Context(), "101").Return(&domain.Account{
		ID:       "101",
		Email:    "test.mail",
		Password: "******",
	}, nil).Once()

	handler.GetAccountDetailsByID(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expectedBody, w.Body.String())

}

func TestHandler_CreateAccount(t *testing.T) {
	data := &domain.Account{
		Email:    "11",
		Password: "******",
	}

	reqBody, _ := json.Marshal(data)

	r := httptest.NewRequest(http.MethodPost, "/account",
		bytes.NewReader(reqBody)).WithContext(context.Background())

	w := httptest.NewRecorder()

	txnService := &MockService{}
	handler := NewHandler(txnService)
	txnService.On("CreateAccount", r.Context(), data).Return(&domain.Account{
		ID:       "101",
		Email:    "test.mail",
		Password: "******",
	}, nil).Once()

	handler.CreateAccount(w, r)

	expectedBody := `{"id": "101","email":"test.mail","password":"******","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expectedBody, w.Body.String())

}
