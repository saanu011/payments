package main

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // postgres
	_ "github.com/golang-migrate/migrate/v4/source/file"       // file migration
	_ "github.com/lib/pq"                                      // psql lib
	"log"
	"payments/config"

	db "payments/pkg/database"
)

func createMigrate(cfg db.Config) (*migrate.Migrate, error) {
	m, err := migrate.New("file://db/migrations", cfg.URL())
	if err != nil {
		return nil, err
	}

	return m, nil
}

func main() {
	cfg, err := config.Load("config.yml")
	if err != nil {
		log.Fatalf("failed to load cofig: %v", err)
	}

	if err := MigrateUp(cfg); err != nil {
		log.Printf("failed to migrate: %v\n rolling back", err)
		if err := RollbackMigration(cfg); err != nil {
			log.Fatalf("failed to rollback: %v", err)
		}
	}
}

func MigrateUp(cfg *config.Config) error {
	m, err := createMigrate(cfg.Database)
	if err != nil {
		return err
	}

	log.Println("[postgresql] starting migration")

	err = m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			log.Printf("[postgresql] no migration performed: %v\n", err)
			return nil
		}

		log.Printf("[postgresql] migration failed: %s\n", err)
		return err
	}

	log.Println("[postgresql] migration successful")

	return nil
}

func RollbackMigration(conf *config.Config) error {
	m, err := createMigrate(conf.Database)
	if err != nil {
		return err
	}

	err = m.Steps(-1)
	if err != nil {
		if err == migrate.ErrNoChange {
			return nil
		}

		return err
	}

	return nil
}
