package config

import (
	"fmt"
	"database/sql"
)

var (
	DB_NAME = ViperGetEnv("DB_NAME")
	DB_USERNAME = ViperGetEnv("DB_USERNAME")
	DB_PASS = ViperGetEnv("DB_PASS")
	DB_HOST = ViperGetEnv("DB_HOST")
	// DB_PORT = ViperGetEnv("DB_PORT")
	DB_SCHEMA = ViperGetEnv("DB_SCHEMA")
)

func SqlConnection() (*sql.DB, error) {
	db, err := sql.Open(fmt.Sprintf("%s", "	%s:%s@tcp(%s:%s)/%s", DB_NAME, DB_USERNAME, DB_PASS, DB_HOST, DB_PORT, DB_SCHEMA))
	if err != nil {
		return nil, err
	}
	return db, nil
}