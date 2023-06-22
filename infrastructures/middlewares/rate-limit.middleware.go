package middlewares

import (
	"Pastebin/infrastructures/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func RateLimit(ac utils.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define a rate limiter with a limit of 10 requests per second
		var limiter = rate.NewLimiter(rate.Every(time.Second), 1)
		// Check if the rate limit has been exceeded
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, "Too many requests")
			return
		}
		c.Next()
	}
}
