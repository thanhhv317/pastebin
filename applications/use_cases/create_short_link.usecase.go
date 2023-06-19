package use_cases

import (
	"Pastebin/infrastructures/orm/mysql/model"
	"Pastebin/infrastructures/repositories"
	"Pastebin/infrastructures/utils"
)

func CreateShortLink(appCtx utils.AppContext, payload model.ShortLink) model.ShortLink {
	db := appCtx.GetMaiDBConnection()

	repositories.CreateShortLink(payload, db)

	return payload
}
