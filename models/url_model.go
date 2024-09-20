package models

import "gorm.io/gorm"

type UrlModel struct {
	gorm.Model
	Id             uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	ShortUrl       string `gorm:"unique" json:"short_url"`
	OriginalString string `json:"original_string"`
}
