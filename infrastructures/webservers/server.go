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
	r.Use(middlewares.Recover(appCtx))

	// import more routes here
	routes.ShortLinkRoute(appCtx, r)

	port := os.Getenv("HTTP_PORT")

	r.Run(":" + port)
}
