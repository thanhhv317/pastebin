package main

import (
	"Pastebin/infrastructures/orm"
	"Pastebin/infrastructures/utils"
	"Pastebin/infrastructures/webservers"
)

func main() {
	db := orm.InitDatabase()
	appCtx := utils.NewAppContext(db)
	
	webservers.InitWebServer(appCtx)
}
