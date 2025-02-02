package store

import (
	"context"
	"database/sql"
	"log"
	"payments/pkg/appError"
	db "payments/pkg/database"
	"payments/src/transaction/domain"
	"strings"
)

type transactionStore struct {
	db db.SQLStore
}

type TransactionStore interface {
	Insert(context.Context, *domain.Transaction) (*domain.Transaction, error)
	GetTransactionByID(context.Context, string) (*domain.Transaction, error)
}

func NewTransactionStore(db db.SQLStore) TransactionStore {
	return &transactionStore{db: db}
}

func (s transactionStore) Insert(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	_, err := s.db.ExecContext(ctx, "INSERT INTO transactions (id, account_id, operation_type_id, amount) "+
		"VALUES ($1, $2, $3, $4)", transaction.ID, transaction.AccountID, int(transaction.OperationTypeID), transaction.Amount)
	if err != nil {
		if strings.Contains(err.Error(), "pq: duplicate key value violates unique constraint") {
			return nil, appError.ErrorDuplicate
		}

		return nil, err
	}

	return transaction, nil
}

func (s transactionStore) GetTransactionByID(ctx context.Context, id string) (*domain.Transaction, error) {
	var data domain.Transaction

	if err := s.db.GetContext(ctx, &data, "SELECT * FROM transactions where id = $1", id); err != nil {
		log.Printf("error while fetching account details with id %s: %v", id, err)

		if err == sql.ErrNoRows {
			return nil, appError.ErrorNotFound
		}

		return nil, appError.ErrorInternal
	}

	return &data, nil
}
