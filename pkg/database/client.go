package database

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"

	// register postgres driver
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

const (
	newrelicPostgresDriver = "nrpostgres"
)

type SqlDB struct {
	*sqlx.DB
}

//type SqlTx struct {
//	*sqlx.Tx
//}

type SQLStore interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, args interface{}) (sql.Result, error)
	Close() error
}

func NewSQLStore(dbConfig Config) (SQLStore, error) {
	log.Printf("connecting to database...")

	db, err := sqlx.Connect(newrelicPostgresDriver, dbConfig.ConnectionString())
	if err != nil {
		return nil, err
	}

	log.Printf("connection to database established")

	sqlDB := &SqlDB{db}

	return sqlDB, nil
}
