package routes

import (
	"Pastebin/infrastructures/utils"
	"Pastebin/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func ShortLinkRoute(appCtx utils.AppContext, r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/short-links", controllers.CreateShortLink(appCtx))
	}
}
