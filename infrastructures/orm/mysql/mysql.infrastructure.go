package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

func InitMysql() (*gorm.DB, *gorm.DB) {

	//dsn := fmt.Sprintf(user + ":" + password + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local")

	masterDsn := os.Getenv("DB_MASTER_URL")
	slaveDsn := os.Getenv("DB_SLAVE_URL")

	dbWrite, errWrite := gorm.Open(mysql.Open(masterDsn), &gorm.Config{})
	if errWrite != nil {
		fmt.Println(errWrite)
	}

	dbRead, errRead := gorm.Open(mysql.Open(slaveDsn), &gorm.Config{})
	if errRead != nil {
		fmt.Println(errRead)
	}

	sqlDBR, errR := dbRead.DB()
	if errR != nil {
		fmt.Println(errR)
	}

	sqlDBW, errW := dbWrite.DB()
	if errW != nil {
		fmt.Println(errW)
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDBR.SetMaxIdleConns(10)
	sqlDBW.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDBR.SetMaxOpenConns(100)
	sqlDBW.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDBR.SetConnMaxLifetime(time.Hour)
	sqlDBW.SetConnMaxLifetime(time.Hour)
	return dbRead, dbWrite
}
