package service_test

import (
	"context"
	"github.com/pkg/errors"
	"payments/src/transaction/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"payments/src/transaction/domain"
	"payments/src/transaction/store"
)

func TestInsertSuccess(t *testing.T) {
	transactionStore := store.NewMockStore(t)
	txnService := service.NewTransactionService(transactionStore)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		txn := &domain.Transaction{
			ID:              "txn-001",
			AccountID:       "101",
			OperationTypeID: domain.Normal,
			Amount:          1000,
		}

		transactionStore.On("Insert", ctx, txn).Return(txn, nil).Once()

		data, err := txnService.CreateTransaction(ctx, txn)
		assert.NoError(t, err)

		assert.NotNil(t, data)
		assert.NotNil(t, data.ID)
		assert.Equal(t, "101", data.AccountID)
		assert.Equal(t, domain.Normal, data.OperationTypeID)
		assert.Equal(t, data.Amount, data.Amount)
	})
}

func TestGetTransactionByIDSuccess(t *testing.T) {
	transactionStore := store.NewMockStore(t)
	txnService := service.NewTransactionService(transactionStore)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		now := time.Now()
		txn := &domain.Transaction{
			ID:              "txn-001",
			AccountID:       "101",
			OperationTypeID: domain.Normal,
			Amount:          1000,
			CreatedAt:       now,
			UpdatedAt:       now,
		}

		transactionStore.On("GetTransactionByID", ctx, txn.ID).Return(txn, nil).Once()

		data, err := txnService.FindTransactionByID(ctx, txn.ID)
		assert.NoError(t, err)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, "101", data.AccountID)
		assert.Equal(t, domain.Normal, data.OperationTypeID)
		assert.Equal(t, data.Amount, data.Amount)
	})
}

func TestGetTransactionByIDNotFoundError(t *testing.T) {
	transactionStore := store.NewMockStore(t)
	txnService := service.NewTransactionService(transactionStore)

	t.Run("transaction not found error", func(t *testing.T) {
		ctx := context.Background()
		txn := &domain.Transaction{
			ID:              "txn-001",
			AccountID:       "101",
			OperationTypeID: domain.Normal,
			Amount:          1000,
		}

		transactionStore.On("GetTransactionByID", ctx, txn.ID).
			Return(nil, errors.New("not found")).Once()

		data, err := txnService.FindTransactionByID(ctx, txn.ID)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "not found")

		assert.Nil(t, data)
	})
}
