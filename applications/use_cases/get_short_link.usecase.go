package use_cases

import (
	"Pastebin/infrastructures/orm/mysql/model"
	"Pastebin/infrastructures/repositories"
	"Pastebin/infrastructures/utils"
)

func GetShortLink(appCtx utils.AppContext, cond map[string]interface{}) (*model.ShortLink, error) {
	db := appCtx.GetMaiDBConnection()
	return repositories.GetShortLink(cond, db)
}
