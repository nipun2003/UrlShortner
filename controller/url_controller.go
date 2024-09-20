package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/nipun2003/url-shortner/services"
)

type UrlController interface {
	HandleCreateShortUrl(c *gin.Context)
	HandleRedirect(c *gin.Context)
}

var urlController UrlController

type UrlControllerImpl struct {
	service services.UrlShortnerService
}

func (ctrl *UrlControllerImpl) HandleCreateShortUrl(c *gin.Context) {
	var url = c.PostForm("url")
	if url == "" {
		c.JSON(400, gin.H{
			"error": "URL is required",
		})
		return
	}
	shortUrl, err := ctrl.service.ShortenUrl(url)
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

func (ctrl *UrlControllerImpl) HandleRedirect(c *gin.Context) {
	shortUrl := c.Param("short_url")
	originalUrl, err := ctrl.service.GetOriginalUrl(shortUrl)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error fetching original URL",
		})
		return
	}
	c.Redirect(301, originalUrl)
}

func NewUrlController() UrlController {
	if urlController == nil {
		urlController = &UrlControllerImpl{
			service: services.NewUrlShortnerService(),
		}
	}
	return urlController
}
