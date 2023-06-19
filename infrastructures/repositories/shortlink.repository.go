package repositories

import (
	"Pastebin/common"
	"Pastebin/infrastructures/orm/mysql/model"
	"gorm.io/gorm"
)

func CreateShortLink(payload model.ShortLink, db *gorm.DB) {
	db.Create(&payload)
}

func GetShortLink(condition map[string]interface{}, db *gorm.DB) (*model.ShortLink, error) {
	var data model.ShortLink
	if err := db.Table(model.ShortLink{}.TableName()).Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
