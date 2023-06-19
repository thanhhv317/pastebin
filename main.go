package main

import (
	"Pastebin/infrastructures/orm/mysql"
	"Pastebin/infrastructures/webservers"
	"fmt"
)

func main() {
	db := mysql.InitMysql()
	webservers.InitWebServer()
	fmt.Println(db)
}
