package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func InitMysql() *gorm.DB {

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf(user + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected successful mysql at:", dsn)

	return db
}
