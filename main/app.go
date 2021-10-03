package main

import (
	"enigmacamp.com/gosql/config"
	"enigmacamp.com/gosql/deliveries"
	"enigmacamp.com/gosql/logger"
	"enigmacamp.com/gosql/manager"
	"enigmacamp.com/gosql/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/http"
)

/*
1. viper config reader => go get github.com/spf13/viper
2. kingpin cli parser => go get gopkg.in/alecthomas/kingpin.v2
3. Buat folder config, config.go
4. Buat folder repositories, lalu db_session_factory.go
5. Buat folder usecase
6. gin/gonic => go get -u github.com/gin-gonic/gin
7. Tambah http config di config.go
8. Buat folder deliveries
9. Buat folder app_response
*/

var (
	appEnvironment = kingpin.Flag("env", "Environment").HintOptions("dev", "prod").Default("dev").Short('e').String()
)

type app struct {
	sf             *repositories.DbSessionFactory
	serviceManager manager.ServiceManager
	config         *config.Config
}

func newApp() app {
	kingpin.Parse()
	c := config.NewConfig(*appEnvironment)
	err := c.InitDb()
	if err != nil {
		logger.Logger.Fatal().Msg(err.Error())
		panic(err)
	}
	myapp := app{
		sf:             c.SessionFactory,
		serviceManager: manager.NewServiceManger(c.SessionFactory),
		config:         c,
	}
	return myapp
}
func (a app) run() {
	appRouter := gin.Default()
	deliveries.NewAppDelivery(appRouter, a.serviceManager).Initialize()
	hostListen := fmt.Sprintf("%v:%v", a.config.HttpConf.Host, a.config.HttpConf.Port)
	logger.Logger.Info().Msgf("Ready to listen on %v", hostListen)

	if err := http.ListenAndServe(hostListen, appRouter); err != nil {
		logger.Logger.Fatal().Msg(err.Error())
	}

	//bulkCustomers, err := a.serviceManager.CustomerUseCase().RegisterBulkCustomer([]models.Customer{
	//	{
	//		FirstName: "Alias",
	//		LastName:  "Rizal",
	//		Address:   "Pondok Kelapa",
	//		City:      "Jakarta",
	//	},
	//	{
	//		FirstName: "Agus",
	//		LastName:  "Riyan",
	//		Address:   "Mangga Lima",
	//		City:      "Jakarta",
	//	},
	//})
	//logger.Logger.Info().Msgf("New Customers : %v \n", bulkCustomers)
	//

	//customers, err := a.serviceManager.CustomerUseCase().FindCustomerByFirstName("ka")
	//logger.Logger.Info().Msgf("Customer : %v \n", customers)
}
func main() {
	newApp().run()
}
