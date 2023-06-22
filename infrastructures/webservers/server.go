package webservers

import (
	"Pastebin/infrastructures/middlewares"
	"Pastebin/infrastructures/utils"
	"Pastebin/interfaces/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func InitWebServer(appCtx utils.AppContext) {
	r := gin.Default()
	r.ForwardedByClientIP = true
	r.Use(middlewares.Recover(appCtx))
	r.Use(middlewares.RateLimit(appCtx))
	// import more routes here
	routes.ShortLinkRoute(appCtx, r)

	port := os.Getenv("HTTP_PORT")

	r.Run(":" + port)
}
