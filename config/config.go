package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
}

func NewFromEnv() (*Config, error) {
	c := &Config{}

	return c, nil
}

type Db struct {
	Username     string
	Password     string
	Host         string
	Port         string
	DatabaseName string
	MaxAttempts  int
	Timeout      int
}

func GetDbConfig() Db {
	var dbConfig Db
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки файла .env")
	}
	dbConfig.Username = os.Getenv("DB_USERNAME")
	dbConfig.Password = os.Getenv("DB_PASSWORD")
	dbConfig.Host = os.Getenv("DB_HOST")
	dbConfig.Port = os.Getenv("DB_PORT")
	dbConfig.DatabaseName = os.Getenv("DB_NAME")
	dbConfig.MaxAttempts = 5
	dbConfig.Timeout = 5
	return dbConfig
}
