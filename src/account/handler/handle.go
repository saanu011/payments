package handler

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
	"payments/pkg/appError"

	"payments/pkg/handler"
	"payments/src/account/domain"
)

type Handler struct {
	service Service
}

type Service interface {
	CreateAccount(context.Context, *domain.Account) (*domain.Account, error)
	FindAccountByID(context.Context, string) (*domain.Account, error)
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var account *domain.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		handler.JSONWithError(w, r, appError.BadRequest(err.Error()))
		return
	}

	newAccount, err := h.service.CreateAccount(ctx, account)
	if errors.Is(err, appError.ErrorDuplicate) {
		handler.JSONWithError(w, r, appError.BadRequest("account already exists"))
		return
	}

	if err != nil {
		handler.JSONWithError(w, r, appError.InternalError(err.Error()))
		return
	}

	handler.JSON(w, r, http.StatusOK, newAccount)

	return
}

func (h Handler) GetAccountDetailsByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id := vars["accountID"]

	data, err := h.service.FindAccountByID(ctx, id)
	if errors.Is(err, appError.ErrorNotFound) {
		handler.JSONWithError(w, r, appError.NotFound(err.Error()))
		return
	}

	if err != nil {
		handler.JSONWithError(w, r, appError.InternalError(err.Error()))
		return
	}

	handler.JSON(w, r, http.StatusOK, data)
	return
}
