package store_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"payments/config"
	"payments/pkg/appError"
	db "payments/pkg/database"
	"payments/src/account/domain"
	"payments/src/account/store"
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
	accountStore := store.NewAccountStore(setup)
	defer TearDownTest(setup)

	t.Run("success", func(t *testing.T) {
		data, err := accountStore.Insert(context.Background(), &domain.Account{
			Email: "test_email", Password: "test_password", ID: "1001"})
		assert.NoError(t, err)

		actual, err := accountStore.GetAccountByID(context.Background(), data.ID)
		assert.NoError(t, err)
		assert.Equal(t, data.Email, actual.Email)
		assert.Equal(t, data.ID, actual.ID)
	})
}

func TestInsertDuplicateError(t *testing.T) {
	setup := SetupTestPrerequisite()
	accountStore := store.NewAccountStore(setup)
	defer TearDownTest(setup)

	t.Run("duplicate key", func(t *testing.T) {
		_, err := accountStore.Insert(context.Background(), &domain.Account{
			Email: "test_email", Password: "test_password", ID: "1001"})
		assert.NoError(t, err)

		_, err = accountStore.Insert(context.Background(), &domain.Account{
			Email: "test_email", Password: "test_password", ID: "1001"})
		assert.Error(t, err)
		assert.Equal(t, appError.ErrorDuplicate, err)
	})
}

func TestGetAccountByIDError(t *testing.T) {
	setup := SetupTestPrerequisite()
	accountStore := store.NewAccountStore(setup)
	defer TearDownTest(setup)

	t.Run("no record found", func(t *testing.T) {
		_, err := accountStore.GetAccountByID(context.Background(), "101")
		assert.Error(t, err)
		assert.Equal(t, appError.ErrorNotFound, err)

	})
}

func TestGetAccountByIDSuccess(t *testing.T) {
	setup := SetupTestPrerequisite()
	accountStore := store.NewAccountStore(setup)
	defer TearDownTest(setup)

	t.Run("success", func(t *testing.T) {
		data, err := accountStore.Insert(context.Background(), &domain.Account{
			Email: "test_email", Password: "test_password", ID: "1001"})
		assert.NoError(t, err)

		actualData, err := accountStore.GetAccountByID(context.Background(), data.ID)
		assert.NoError(t, err)
		assert.Equal(t, data.ID, actualData.ID)
		assert.Equal(t, data.Email, actualData.Email)
		assert.Equal(t, data.Password, actualData.Password)
	})
}
