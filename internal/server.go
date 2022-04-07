package internal

import (
	"fmt"
	"github.com/fades-io/reg/internal/logs"
	"log"
	"os"

	"github.com/fades-io/reg/internal/db/postgresql"
	"github.com/fades-io/reg/internal/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var srv = server.Server{}

// DbConfig Конфиг БД
type DbConfig struct {
	driver   string
	host     string
	port     string
	name     string
	user     string
	password string
}

// Run Создание и запуск сервера
func Run() {
	dbConfig := GetDbConfig()
	storage := GetDB(dbConfig)
	srv.Init(storage)
	srv.Run()
}

// GetDbConfig Получение конфигурации для БД
func GetDbConfig() *DbConfig {
	return &DbConfig{
		driver:   os.Getenv("DB_DRIVER"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		name:     os.Getenv("POSTGRES_DB"),
		user:     os.Getenv("POSTGRES_USER"),
		password: os.Getenv("POSTGRES_PASSWORD"),
	}
}

// GetDB Получение БД
func GetDB(dbConfig *DbConfig) server.Storage {
	if dbConfig.driver == "postgres" {
		dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbConfig.host, dbConfig.port, dbConfig.name, dbConfig.user, dbConfig.password)
		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Printf(logs.DatabaseAccessDenied, dbConfig.driver)
			log.Fatal(logs.Error, err)
		} else {
			fmt.Printf(logs.DatabaseConnection, dbConfig.driver)
		}
		return postgresql.NewDb(gormDB)
	}
	return nil
}
