package deliveries

import (
	"enigmacamp.com/gosql/appresponse"
	"enigmacamp.com/gosql/models"
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	prefix string
	rt     *gin.Engine
}

func (cr *AuthRoute) InitRoute() {
	vCustomer := cr.rt.Group(cr.prefix)
	vCustomer.POST("", cr.userAuth)
}

func (cr *AuthRoute) userAuth(c *gin.Context) {
	var login models.UserCred
	err := c.BindJSON(&login)
	if err != nil {
		appresponse.NewJsonResponse(c).SendError(appresponse.NewUnauthorizedError(err))
	}
	c.SetCookie("appsession", "123", 10, "/", "localhost", false, true)
	appresponse.NewJsonResponse(c).SendData(appresponse.NewSimpleResponseMessage("SUCCESS"))
}

func NewAuthRoute(prefix string, rt *gin.Engine) IAppRouter {
	return &AuthRoute{
		prefix,
		rt,
	}
}
