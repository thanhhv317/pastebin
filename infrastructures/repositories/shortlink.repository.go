package repositories

import (
	"Pastebin/infrastructures/orm/mysql/model"
	"gorm.io/gorm"
)

func CreateShortLink(payload model.ShortLink, db *gorm.DB) {
	db.Create(&payload)
}
