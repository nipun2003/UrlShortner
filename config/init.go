package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var dbHost string
var dbPort string
var dbUser string
var dbPass string
var dbName string
var dbSchema string

func InitializeEnv() {
	// Load envs
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
	dbSchema = os.Getenv("DB_SCHEMA")

}

func GetDBUrl() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)
}

func GetSchema() string {
	return dbSchema
}

func GetDBUrlWithSchema() string {
	if dbSchema == "" {
		dbSchema = "public"
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName, dbSchema)
}
