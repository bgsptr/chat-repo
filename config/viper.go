package config 

import (
	"github.com/spf13/viper"
	"log"
)

func ViperGetEnv(key string) string {
	viper.SetConfigType("env")
	viper.SetConfigName("")
	viper.AddConfigPath("D:/DEVELOPMENT/golang/src/group_service")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Can't find the env file")
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion for key '%s'", key)
	}

	return value
}