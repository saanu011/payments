package store_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"payments/config"
	"payments/pkg/appError"
	db "payments/pkg/database"
	accountDomain "payments/src/account/domain"
	accountStore "payments/src/account/store"
	"payments/src/transaction/domain"
	"payments/src/transaction/store"
)

func SetupTestPrerequisite() db.SQLStore {
	conf, err := config.Load("../../../config.yml")
	if err != nil {
		log.Fatalf("Error loading configuration file: %v", err)
	}

	st, err := db.NewSQLStore(conf.Database)
	if err != nil {
		log.Fatalf("Error creating store: %v", err)
	}

	return st
}

func TearDownTest(s db.SQLStore) {
	_, err := s.ExecContext(context.Background(), "TRUNCATE transactions, accounts")
	if err != nil {
		log.Fatalf("Error truncating table: %v", err)
	}
}

func TestInsertSuccess(t *testing.T) {
	setup := SetupTestPrerequisite()
	transactionStore := store.NewTransactionStore(setup)
	accStore := accountStore.NewAccountStore(setup)
	accountData, err := accStore.Insert(context.Background(), &accountDomain.Account{
		Email: "test_email", Password: "test_password", ID: "1001"})
	assert.NoError(t, err)

	defer TearDownTest(setup)

	t.Run("success", func(t *testing.T) {
		data, err := transactionStore.Insert(context.Background(), &domain.Transaction{
			ID: "txn-001", AccountID: accountData.ID, OperationTypeID: domain.Normal, Amount: 1000})
		assert.NoError(t, err)

		actual, err := transactionStore.GetTransactionByID(context.Background(), data.ID)
		assert.NoError(t, err)
		assert.NotNil(t, actual)
		assert.Equal(t, accountData.ID, actual.AccountID)
		assert.Equal(t, domain.Normal, actual.OperationTypeID)
		assert.Equal(t, data.Amount, actual.Amount)
	})
}

func TestGetTransactionByIDError(t *testing.T) {
	setup := SetupTestPrerequisite()
	transactionStore := store.NewTransactionStore(setup)
	defer TearDownTest(setup)

	t.Run("no record found", func(t *testing.T) {
		_, err := transactionStore.GetTransactionByID(context.Background(), "101")
		assert.Error(t, err)
		assert.Equal(t, appError.ErrorNotFound, err)
	})
}

func TestGetTransactionByIDSuccess(t *testing.T) {
	setup := SetupTestPrerequisite()
	transactionStore := store.NewTransactionStore(setup)
	accStore := accountStore.NewAccountStore(setup)
	accountData, err := accStore.Insert(context.Background(), &accountDomain.Account{
		Email: "test_email", Password: "test_password", ID: "101"})
	assert.NoError(t, err)
	defer TearDownTest(setup)

	t.Run("success", func(t *testing.T) {
		data, err := transactionStore.Insert(context.Background(), &domain.Transaction{
			AccountID: accountData.ID, OperationTypeID: domain.Withdrawal, Amount: -1000, ID: "1001"})
		assert.NoError(t, err)

		actualData, err := transactionStore.GetTransactionByID(context.Background(), data.ID)
		assert.NoError(t, err)
		assert.Equal(t, data.ID, actualData.ID)
		assert.Equal(t, data.AccountID, actualData.AccountID)
		assert.Equal(t, data.OperationTypeID, actualData.OperationTypeID)
		assert.Equal(t, data.Amount, actualData.Amount)
	})
}
