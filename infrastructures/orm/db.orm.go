package orm

import (
	"Pastebin/infrastructures/orm/mysql"
	"gorm.io/gorm"
)

type Any int

func InitDatabase() *gorm.DB {
	db := mysql.InitMysql()

	return db
}
