package deliveries

import (
	"enigmacamp.com/gosql/appresponse"
	"enigmacamp.com/gosql/manager"
	"enigmacamp.com/gosql/models"
	"enigmacamp.com/gosql/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

type CustomerRoute struct {
	prefix  string
	useCase usecases.ICustomerUseCase
	rt      *gin.Engine
}

func (cr *CustomerRoute) InitRoute() {
	vCustomer := cr.rt.Group(cr.prefix)
	//vCustomer.Use(authMiddleware())
	vCustomer.Use(authCookieMiddleware())
	vCustomer.GET("/total", cr.getTotalCustomer)
	vCustomer.GET("", cr.findCustomer)
	vCustomer.POST("", cr.registerCustomer)
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

func (cr *CustomerRoute) registerCustomer(c *gin.Context) {
	firstName := c.PostForm("firstName")
	lastName := c.PostForm("lastName")
	address := c.PostForm("address")
	city := c.PostForm("city")

	file, header, err := c.Request.FormFile("avatar")
	newCustomer := models.Customer{
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		City:      city,
	}
	customer, err := cr.useCase.RegisterNewCustomer(newCustomer)
	if err != nil {
		appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err))
	}

	filename := fmt.Sprintf("%s%s", customer.Id, filepath.Ext(filepath.Ext(header.Filename)))
	err = cr.useCase.UploadAvatar(filename, file)
	if err != nil {
		appresponse.NewJsonResponse(c).SendError(appresponse.NewInternalServerError(err))
	}
	appresponse.NewJsonResponse(c).SendData(appresponse.NewSimpleResponseMessage(customer))
}

func NewCustomerRoute(prefix string, sf manager.ServiceManager, rt *gin.Engine) IAppRouter {
	return &CustomerRoute{
		prefix,
		sf.CustomerUseCase(),
		rt,
	}
}
