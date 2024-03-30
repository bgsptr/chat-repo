package config

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	DB_SCHEMA = os.Getenv("DB_SCHEMA")
	DB_MIGRATE_LOCATION = os.Getenv("DB_MIGRATE_LOCATION")
)

func RunPostgresMigrate(DB *gorm.DB) {
    if DB_SCHEMA == "" {
        log.Println("DB_SCHEMA environment variable not set")
        return
    }

    if DB_MIGRATE_LOCATION == "" {
        log.Println("DB_MIGRATE_LOCATION environment variable not set")
        return
    }

    sql, errdb := DB.DB()
    if errdb != nil {
        log.Println("error fetching SQL connection:", errdb)
        return
    }

    driver, err := postgres.WithInstance(sql, &postgres.Config{})
    if err != nil {
        log.Println("error creating driver instance:", err)
        return
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://" + DB_MIGRATE_LOCATION,
        DB_SCHEMA,
        driver,
    )
    if err != nil {
        log.Println("error creating migration instance:", err)
        return
    }

    if err := m.Steps(2); err != nil {
        log.Println("error executing migration steps:", err)
        return
    }

    log.Println("database migration successful")
}
