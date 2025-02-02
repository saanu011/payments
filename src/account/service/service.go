package service

import (
	"context"
	"github.com/google/uuid"

	"payments/src/account/domain"
	"payments/src/account/store"
)

type Service struct {
	store store.AccountStore
}

func NewAccountService(store store.AccountStore) *Service {
	return &Service{store: store}
}

func (s *Service) CreateAccount(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	account.ID = uuid.New().String()

	data, err := s.store.Insert(ctx, account)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *Service) FindAccountByID(ctx context.Context, id string) (*domain.Account, error) {
	return s.store.GetAccountByID(ctx, id)
}
