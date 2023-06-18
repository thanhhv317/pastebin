package routes

import (
	"Pastebin/interfaces/controllers"
	"github.com/gin-gonic/gin"
)

func ShortLinkRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.POST("/short-links", controllers.CreateShortLink())
	}
}
