package config

import (
	"database/sql"
	// "errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)
   
func RunPostgresMigrate(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
	 log.Println(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
	 "file://./migrations/",
	 "whatsapp", driver)
	if err != nil {
	 log.Println(err)
	}
	if err := m.Steps(2); err != nil {
	 log.Println(err)
	}
}