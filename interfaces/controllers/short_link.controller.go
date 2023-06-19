package controllers

import (
	"Pastebin/applications/use_cases"
	"Pastebin/common"
	"Pastebin/infrastructures/orm/mysql/model"
	"Pastebin/infrastructures/utils"
	"Pastebin/interfaces/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateShortLink(appCtx utils.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data dtos.CreateShortLinkDto
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)

		var shortLink model.ShortLink = model.ShortLink{
			Link:      data.Link,
			Shortlink: utils.Base62Encode(uint64(timestamp)),
			ExpiredAt: time.Now().AddDate(0, 0, 10),
		}

		// call usecase to create shortlink
		use_cases.CreateShortLink(appCtx, shortLink)

		c.JSON(http.StatusOK, gin.H{
			"data": shortLink,
		})
	}
}

func GetShortLink(appCtx utils.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var data dtos.GetShortLinkDto
		if err := c.ShouldBindUri(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		result, err := use_cases.GetShortLink(appCtx, map[string]interface{}{"shortlink": data.ShortLink})
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
