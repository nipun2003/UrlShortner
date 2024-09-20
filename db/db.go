package db

import (
	"fmt"
	"log/slog"

	"github.com/nipun2003/url-shortner/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Generate a Final tag
const TAG = "DB"

var DB *gorm.DB = nil

func InitDb() {
	var err error
	DB, err = gorm.Open(postgres.Open(config.GetDBUrl()), &gorm.Config{})
	slog.Info(TAG, "URL", config.GetDBUrl())
	if err != nil {
		panic(err)
	}
	// Create the schema if it does not exist
	err = DB.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", config.GetSchema())).Error
	if err != nil {
		panic("failed to create schema")
	}
	DB, err = gorm.Open(postgres.Open(config.GetDBUrlWithSchema()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database with schema")
	}
	slog.Info("Connected to database")
}
