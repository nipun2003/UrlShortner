package services

import (
	"log/slog"

	"github.com/nipun2003/url-shortner/db"
	"github.com/nipun2003/url-shortner/models"
	"github.com/nipun2003/url-shortner/utils"
	"gorm.io/gorm"
)

var _instance UrlShortnerService

type UrlShortnerService interface {
	ShortenUrl(url string) (string, error)
	GetOriginalUrl(shortUrl string) (string, error)
}

type UrlShortnerServiceImpl struct {
	shortIDService ShortIDGenerateService
	rdb            *db.RedisClient
	db             *gorm.DB
}

func NewUrlShortnerService() UrlShortnerService {
	if _instance == nil {
		_instance = &UrlShortnerServiceImpl{
			shortIDService: NewShortIDGenerateService(),
			rdb:            db.NewRedisClient(),
			db:             db.DB,
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
	// Save the short url in db first
	err = u.db.Create(&models.UrlModel{
		ShortUrl:       shortUrl,
		OriginalString: url,
	}).Error
	if err != nil {
		return "", err
	}
	// Save the short url in redis
	err = u.rdb.Client.Set(u.rdb.Ctx, shortUrl, url, 0).Err()
	if err != nil {
		return "", err
	}
	return shortUrl, nil
}

func (u *UrlShortnerServiceImpl) GetOriginalUrl(shortUrl string) (string, error) {
	// GetOriginalUrl function
	slog.Info("GetOriginalURL", "Getting Original URL", shortUrl)
	// Check in redis first
	url, err := u.rdb.Client.Get(u.rdb.Ctx, shortUrl).Result()
	if err == nil {
		return url, nil
	}
	var urlModel models.UrlModel
	err = u.db.Where("short_url = ?", shortUrl).First(&urlModel).Error
	if err != nil {
		return "", err
	}
	// Save the short url in redis
	err = u.rdb.Client.Set(u.rdb.Ctx, shortUrl, urlModel.OriginalString, 0).Err()
	if err != nil {
		return "", err
	}
	return urlModel.OriginalString, nil
}
