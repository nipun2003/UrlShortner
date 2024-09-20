package db

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/nipun2003/url-shortner/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Generate a Final tag
const TAG = "DB"

var DB *gorm.DB = nil

var dbHost string
var dbPort string
var dbUser string
var dbPass string
var dbName string
var dbSchema string

func InitDb() {
	initValues()
	var err error
	DB, err = gorm.Open(postgres.Open(getDBUrl()), &gorm.Config{})
	slog.Info(TAG, "URL", getDBUrl())
	if err != nil {
		panic(err)
	}
	// Create the schema if it does not exist
	err = DB.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", dbSchema)).Error
	if err != nil {
		panic("failed to create schema")
	}
	DB, err = gorm.Open(postgres.Open(getDBUrlWithSchema()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database with schema")
	}
	slog.Info("Connected to database")
	migrateDb()
}

func initValues() {
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
	dbSchema = os.Getenv("DB_SCHEMA")
}

func getDBUrl() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
}

func GetSchema() string {
	return dbSchema
}

func getDBUrlWithSchema() string {
	if dbSchema == "" {
		dbSchema = "public"
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName, dbSchema)
}

func migrateDb() {
	DB.AutoMigrate(&models.UrlModel{})
}
