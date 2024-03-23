package usersqldb

import (
	// "database/sql"
	// _ "github.com/go-sql-driver/mysql"
	// "os"
	"fmt"
	// "strconv"
	// "log"
	"os"

	// "github.com/spf13/viper"
	"gorm.io/gorm"

	// "gorm.io/driver/mysql"
	"errors"

	"gorm.io/driver/postgres"
)

// func viperEnvVar(key string) string {
// 	viper.SetConfigType("env")
// 	viper.SetConfigName("app")
// 	viper.AddConfigPath("D:/DEVELOPMENT/golang/src/user_service")

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Println("Can't find the env file")
// 	}

// 	value, ok := viper.Get(key).(string)
// 	if !ok {
// 		log.Fatalf("Invalid type assertion for key '%s'", key)
// 	}

// 	return value
// }

var (
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	// DB_PORT = os.Getenv("DB_PORT")
	DB_SCHEMA = os.Getenv("DB_SCHEMA")
)


func NewGorm() (*gorm.DB, error) {
	 // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USERNAME, DB_PASS, DB_HOST, DB_PORT, DB_SCHEMA)

	// dbPort, err := strconv.ParseUint(DB_PORT, 10, 32)
	// if err != nil {
	// 	// Handle the error, e.g., log it and return an error
	// 	return nil, err
	// }
	dbPort := 5432

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", DB_HOST, DB_USERNAME, DB_PASS, DB_SCHEMA, dbPort)
	 db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	 if err != nil {
		return nil, errors.New("database not found")
	 }
	 return db, nil
}