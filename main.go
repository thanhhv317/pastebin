package main

import (
	"Pastebin/infrastructures/orm"
	"Pastebin/infrastructures/utils"
	"Pastebin/infrastructures/webservers"
)

func main() {
	dbRead, dbWrite := orm.InitDatabase()
	appCtx := utils.NewAppContext(dbRead, dbWrite)

	webservers.InitWebServer(appCtx)
}
