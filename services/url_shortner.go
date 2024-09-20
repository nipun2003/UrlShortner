package services

import (
	"log/slog"

	"github.com/nipun2003/url-shortner/utils"
)

var _instance UrlShortnerService

type UrlShortnerService interface {
	ShortenUrl(url string) (string, error)
}

type UrlShortnerServiceImpl struct {
	shortIDService ShortIDGenerateService
}

func NewUrlShortnerService() UrlShortnerService {
	if _instance == nil {
		_instance = &UrlShortnerServiceImpl{
			shortIDService: NewShortIDGenerateService(),
		}
	}
	return _instance
}

func (u *UrlShortnerServiceImpl) ShortenUrl(url string) (string, error) {
	// ShortenUrl function
	slog.Info("ShortenURL", "Shortening URL", url)
	id, err := u.shortIDService.GenerateUniqueID()
	if err != nil {
		return "", err
	}
	shortUrl := utils.Base62Encode(id)
	return shortUrl, nil
}
