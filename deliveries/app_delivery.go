package deliveries

import (
	"enigmacamp.com/gosql/manager"
	"github.com/gin-gonic/gin"
)

const (
	CustomerMainRoute = "/customer"
)

type IAppRouter interface {
	InitRoute()
}
type AppDelivery struct {
	rt *gin.Engine
	sm manager.ServiceManager
}

func NewAppDelivery(rt *gin.Engine, sm manager.ServiceManager) *AppDelivery {
	return &AppDelivery{rt, sm}
}

func (a *AppDelivery) Initialize() {
	routerList := []IAppRouter{
		NewCustomerRoute(CustomerMainRoute, a.sm, a.rt),
	}
	for _, rt := range routerList {
		rt.InitRoute()
	}
}
