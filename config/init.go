package config

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/nipun2003/url-shortner/db"
)

func InitializeApp() {
	// Load envs
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}
	db.InitDb()
	db.InitZookeeper()
	db.InitRedis()
}
