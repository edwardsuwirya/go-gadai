package deliveries

import (
	"enigmacamp.com/gosql/appresponse"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")
		strArr := strings.Split(bearerToken, " ")
		if len(strArr) == 2 {
			token := strArr[1]
			if token == "123" {
				c.Next()
			} else {
				appresponse.NewJsonResponse(c).SendError(appresponse.NewUnauthorizedError(errors.New("Invalid Token")))
			}
		} else {
			appresponse.NewJsonResponse(c).SendError(appresponse.NewUnauthorizedError(errors.New("Invalid Token")))
		}
	}
}
