package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nipun2003/url-shortner/services"
)

type UrlController interface {
	HandleCreateShortUrl(c *gin.Context)
}

var urlController UrlController

type UrlControllerImpl struct {
	urlService services.UrlShortnerService
}

func (controller *UrlControllerImpl) HandleCreateShortUrl(c *gin.Context) {
	var url = c.PostForm("url")
	if url == "" {
		c.JSON(400, gin.H{
			"error": "URL is required",
		})
		return
	}
	shortUrl, err := controller.urlService.ShortenUrl(url)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error generating short URL",
		})
		return
	}

	c.JSON(200, gin.H{
		"short_url": shortUrl,
	})
}

func NewUrlController() UrlController {
	if urlController == nil {
		urlController = &UrlControllerImpl{
			urlService: services.NewUrlShortnerService(),
		}
	}
	return urlController
}
