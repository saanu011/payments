package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"payments/pkg/appError"
	"payments/pkg/handler"
	"payments/src/transaction/domain"
)

type Handler struct {
	service Service
}

type Service interface {
	CreateTransaction(context.Context, *domain.Transaction) (*domain.Transaction, error)
	FindTransactionByID(context.Context, string) (*domain.Transaction, error)
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h *Handler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var transaction *domain.Transaction

	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		handler.JSONWithError(w, r, appError.BadRequest(err.Error()))
		return
	}

	newTxn, err := h.service.CreateTransaction(ctx, transaction)
	if err != nil {
		handler.JSONWithError(w, r, appError.InternalError(err.Error()))
		return
	}

	handler.JSON(w, r, http.StatusOK, newTxn)

	return
}

func (h *Handler) GetTransactionDetailsByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id := vars["transactionID"]

	data, err := h.service.FindTransactionByID(ctx, id)
	if err != nil {
		handler.JSONWithError(w, r, appError.InternalError(err.Error()))
		return
	}

	handler.JSON(w, r, http.StatusOK, data)

	return
}
