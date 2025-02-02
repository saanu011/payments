package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"payments/src/account/handler"
	txnHandler "payments/src/transaction/handler"
)

func NewRouter(deps *Dependencies) *mux.Router {
	r := mux.NewRouter()

	// here goes code to register routes
	r.HandleFunc("/ping", PingHandler).Methods(http.MethodGet)

	accountHandler := handler.NewHandler(deps.AccountService)
	r.HandleFunc("/account", accountHandler.CreateAccount).Methods(http.MethodPost)
	r.HandleFunc("/account/{accountID}", accountHandler.GetAccountDetailsByID).Methods(http.MethodGet)

	transactionHandler := txnHandler.NewHandler(deps.TransactionService)
	r.HandleFunc("/transaction", transactionHandler.CreateTransaction).Methods(http.MethodPost)
	r.HandleFunc("/transaction/{transactionID}", transactionHandler.GetTransactionDetailsByID).Methods(http.MethodGet)

	return r
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte(`pong`)); err != nil {
		log.Printf("Ping Error %v\n", err)
	}
}
