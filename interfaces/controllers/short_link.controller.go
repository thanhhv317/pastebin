package controllers

import (
	"Pastebin/infrastructures/orm/mysql/model"
	"Pastebin/infrastructures/utils"
	"Pastebin/interfaces/dtos"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateShortLink() func(ctx *gin.Context) {
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

		fmt.Println(shortLink)
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}

func GetShortLink() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	}
}
