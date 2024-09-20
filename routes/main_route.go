package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nipun2003/url-shortner/controller"
)

func CreateMainRoutes(router *gin.Engine) {
	urlController := controller.NewUrlController()
	var api = router.Group("/api")
	{
		api.POST("/shorten", urlController.HandleCreateShortUrl)
	}
}
