package webservers

import (
	"Pastebin/interfaces/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func InitWebServer() {
	r := gin.Default()

	// import more routes here
	routes.ShortLinkRoute(r)

	port := os.Getenv("HTTP_PORT")

	r.Run(":" + port)
}
