package deliveries

import (
	"enigmacamp.com/gosql/appresponse"
	"enigmacamp.com/gosql/manager"
	"enigmacamp.com/gosql/usecases"
	"github.com/gin-gonic/gin"
)

type CustomerRoute struct {
	prefix  string
	useCase usecases.ICustomerUseCase
	rt      *gin.Engine
}

func (cr *CustomerRoute) InitRoute() {
	vCustomer := cr.rt.Group(cr.prefix)
	vCustomer.Use(authMiddleware())
	vCustomer.GET("/total", cr.getTotalCustomer)
	vCustomer.GET("", cr.findCustomer)
}

func (cr *CustomerRoute) getTotalCustomer(c *gin.Context) {
	customer, err := cr.useCase.GetTotalCustomer()
	if err != nil {
		appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err))
	}
	appresponse.NewJsonResponse(c).SendData(appresponse.NewSimpleResponseMessage(customer))
}

func (cr *CustomerRoute) findCustomer(c *gin.Context) {
	firstName := c.DefaultQuery("firstName", "")
	if firstName != "" {
		customer, err := cr.useCase.FindCustomerByFirstName(firstName)
		if err != nil {
			appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err))
		}
		appresponse.NewJsonResponse(c).SendData(appresponse.NewSimpleResponseMessage(customer))
	} else {
		appresponse.NewJsonResponse(c).SendData(appresponse.NewResponseMessage("No Data Found", nil))
	}

}

func NewCustomerRoute(prefix string, sf manager.ServiceManager, rt *gin.Engine) IAppRouter {
	return &CustomerRoute{
		prefix,
		sf.CustomerUseCase(),
		rt,
	}
}
