package app

import (
	"payments/config"
	db "payments/pkg/database"
	accountService "payments/src/account/service"
	accountStore "payments/src/account/store"
	transactionService "payments/src/transaction/service"
	transactionStore "payments/src/transaction/store"
)

type Dependencies struct {
	AccountService     *accountService.Service
	TransactionService *transactionService.Service
}

func NewDependencies(conf *config.Config) (*Dependencies, error) {
	sqlStore, err := db.NewSQLStore(conf.Database)
	if err != nil {
		return nil, err
	}

	account := accountService.NewAccountService(accountStore.NewAccountStore(sqlStore))
	transaction := transactionService.NewTransactionService(transactionStore.NewTransactionStore(sqlStore))

	return &Dependencies{
		AccountService:     account,
		TransactionService: transaction,
	}, nil
}
