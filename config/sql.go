package config

import (
	"fmt"
	"database/sql"g
)

func ViperGetEnv(key string) string {
	viper.SetConfigType("env")
	viper.SetConfigName("")
	viper.AddConfigPath("D:/DEVELOPMENT/golang/src/group_service")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Can't find the env file")
	}

	value, ok := viper.Get(key).string
	if !ok {
		log.Fatalf("Invalid type assertion for key '%s'", key)
	}

	return value
}

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