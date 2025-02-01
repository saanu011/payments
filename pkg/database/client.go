package database

import (
	"github.com/jmoiron/sqlx"

	// register postgres driver
	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

const (
	newrelicPostgresDriver = "nrpostgres"
)

type SqlDB struct {
	*sqlx.DB
}

type SqlTx struct {
	*sqlx.Tx
}

func NewSQLStore(dbConfig Config) (*SqlDB, error) {
	db, err := sqlx.Connect(newrelicPostgresDriver, dbConfig.ConnectionString())
	if err != nil {
		return nil, err
	}

	sqlDB := &SqlDB{db}

	return sqlDB, nil
}
