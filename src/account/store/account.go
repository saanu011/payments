package store

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"payments/pkg/appError"
	db "payments/pkg/database"
	"payments/src/account/domain"
)

type accountStore struct {
	db db.SQLStore
}

type AccountStore interface {
	Insert(context.Context, *domain.Account) (*domain.Account, error)
	GetAccountByID(context.Context, string) (*domain.Account, error)
}

func NewAccountStore(db db.SQLStore) AccountStore {
	return &accountStore{db: db}
}

func (s accountStore) Insert(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	_, err := s.db.ExecContext(ctx, "INSERT INTO accounts (id, email, password) "+
		"VALUES ($1, $2, $3)", account.ID, account.Email, account.Password)
	if err != nil {
		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			return nil, appError.ErrorDuplicate
		}

		return nil, err
	}

	return account, nil
}

func (s accountStore) GetAccountByID(ctx context.Context, id string) (*domain.Account, error) {
	var data domain.Account

	if err := s.db.GetContext(ctx, &data, "SELECT * FROM accounts WHERE id = $1", id); err != nil {
		log.Printf("error while fetching account details with id %s: %v", id, err)

		if err == sql.ErrNoRows {
			return nil, appError.ErrorNotFound
		}

		return nil, appError.ErrorInternal
	}

	return &data, nil
}

func (s accountStore) GetAccountByEmail(ctx context.Context, email string) (*domain.Account, error) {
	var data domain.Account

	if err := s.db.GetContext(ctx, &data, "SELECT * FROM accounts WHERE email = $1", email); err != nil {
		log.Printf("error while fetching account details with email %s: %v", email, err)

		if err == sql.ErrNoRows {
			return nil, appError.ErrorNotFound
		}

		return nil, appError.ErrorInternal
	}

	return &data, nil
}
