package deliveries

import (
	"enigmacamp.com/gosql/appresponse"
	"errors"
	"github.com/gin-gonic/gin"
)

func authCookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("appsession")
		if cookie == "123" {
			c.Next()
		} else {
			appresponse.NewJsonResponse(c).SendError(appresponse.NewUnauthorizedError(errors.New("Invalid session")))
		}
	}
}
