package orm

import (
	"Pastebin/infrastructures/orm/mysql"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, *gorm.DB) {
	dbRead, dbWrite := mysql.InitMysql()

	return dbRead, dbWrite
}
