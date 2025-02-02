package service

import (
	"context"

	"github.com/google/uuid"

	"payments/src/transaction/domain"
	"payments/src/transaction/store"
)

type Service struct {
	store store.TransactionStore
}

func NewTransactionService(store store.TransactionStore) *Service {
	return &Service{store: store}
}

func (s *Service) CreateTransaction(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	if transaction.OperationTypeID == domain.Withdrawal || transaction.OperationTypeID == domain.CreditVoucher {
		transaction.Amount = -transaction.Amount
	}

	transaction.ID = uuid.New().String()

	return s.store.Insert(ctx, transaction)
}

func (s *Service) FindTransactionByID(ctx context.Context, id string) (*domain.Transaction, error) {
	return s.store.GetTransactionByID(ctx, id)
}
