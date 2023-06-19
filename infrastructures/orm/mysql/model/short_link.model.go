package model

import "time"

const EntityName = "ShortLink"

type ShortLink struct {
	Shortlink string    `json:"shortlink" gorm:"column:shortlink";`
	Link      string    `json:"link" gorm:"column:link";`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at";`
}

func (ShortLink) TableName() string {
	return "shortlinks"
}
