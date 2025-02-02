package service_test

import (
	"context"
	"github.com/pkg/errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"payments/src/account/domain"
	"payments/src/account/service"
	"payments/src/account/store"
)

func TestInsertSuccess(t *testing.T) {
	accountStore := store.NewMockStore(t)
	accountService := service.NewAccountService(accountStore)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		txn := &domain.Account{
			ID:       "-001",
			Email:    "mail.com",
			Password: "******",
		}

		accountStore.On("Insert", ctx, txn).Return(txn, nil).Once()

		data, err := accountService.CreateAccount(ctx, txn)
		assert.NoError(t, err)

		assert.NotNil(t, data)
		assert.NotNil(t, data.ID)
		assert.Equal(t, "mail.com", data.Email)
	})
}

func TestGetAccountByIDSuccess(t *testing.T) {
	accountStore := store.NewMockStore(t)
	accountService := service.NewAccountService(accountStore)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		now := time.Now()
		txn := &domain.Account{
			ID:        "-001",
			Email:     "mail.com",
			Password:  "******",
			CreatedAt: now,
			UpdatedAt: now,
		}

		accountStore.On("GetAccountByID", ctx, txn.ID).Return(txn, nil).Once()

		data, err := accountService.FindAccountByID(ctx, txn.ID)
		assert.NoError(t, err)

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, "mail.com", data.Email)
	})
}

func TestGetAccountByIDNotFoundError(t *testing.T) {
	accountStore := store.NewMockStore(t)
	accountService := service.NewAccountService(accountStore)

	t.Run("account not found error", func(t *testing.T) {
		ctx := context.Background()
		now := time.Now()
		txn := &domain.Account{
			ID:        "-001",
			Email:     "mail.com",
			Password:  "******",
			CreatedAt: now,
			UpdatedAt: now,
		}

		accountStore.On("GetAccountByID", ctx, txn.ID).
			Return(nil, errors.New("not found")).Once()

		data, err := accountService.FindAccountByID(ctx, txn.ID)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "not found")

		assert.Nil(t, data)
	})
}
