package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/nipun2003/url-shortner/config"
	"github.com/nipun2003/url-shortner/db"
	"github.com/nipun2003/url-shortner/routes"
)

func init() {
	config.InitializeApp()
}

func main() {
	slog.Info("---------------------Starting Server---------------------")
	var router = gin.Default()
	routes.CreateMainRoutes(router)
	err := router.Run()
	if err != nil {
		slog.Error(err.Error())
	}
	db.CloseZookeeperConnection()
}
