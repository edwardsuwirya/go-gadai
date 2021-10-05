package config

import (
	appLog "enigmacamp.com/gosql/logger"
	repo "enigmacamp.com/gosql/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

type dbConf struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	schema     string
	dbEngine   string
}

type HttpConf struct {
	HttpServe      string
	maxFileSize    int64
	UploadFilePath string
}

type Config struct {
	SessionFactory *repo.DbSessionFactory
	dbConf         *dbConf
	HttpConf       *HttpConf
	Router         *gin.Engine
}

func NewConfig(env string) *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Can not read config")
	}

	outputWriter := os.Stdout
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if env == "dev" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger := zerolog.New(outputWriter).With().Logger()

	if env != "prod" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	c := new(Config)
	c.dbConf = &dbConf{
		dbUser:     c.GetEnv("dbuser", "root"),
		dbPassword: c.GetEnv("dbpassword", "P@ssw0rd"),
		dbHost:     c.GetEnv("dbhost", "localhost"),
		dbPort:     c.GetEnv("dbport", "3306"),
		schema:     c.GetEnv("dbschema", "enigma"),
		dbEngine:   c.GetEnv("dbengine", "mysql"),
	}
	host := c.GetEnv("httphost", "localhost")
	port := c.GetEnv("httpport", "8080")
	fileSize, _ := strconv.ParseInt(c.GetEnv("maxfilesize", "1"), 10, 64)
	uploadFilePath := c.GetEnv("filePath", "public")
	c.HttpConf = &HttpConf{
		HttpServe:      fmt.Sprintf("%v:%v", host, port),
		maxFileSize:    fileSize,
		UploadFilePath: uploadFilePath,
	}
	appLog.Logger = logger
	return c
}

func (c *Config) InitDb() error {
	appLog.Logger.Debug().Msg("======= Create DB Connection =======")
	sf := repo.NewDbSessionFactory(c.dbConf.dbEngine, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.dbConf.dbUser, c.dbConf.dbPassword, c.dbConf.dbHost, c.dbConf.dbPort, c.dbConf.schema))
	c.SessionFactory = sf
	return nil
}

func (c *Config) InitRouter() {
	appLog.Logger.Debug().Msg("======= Create Http Router =======")
	appRouter := gin.Default()
	// "8 times 2, 30 times" = 8388608 byte
	appRouter.MaxMultipartMemory = c.HttpConf.maxFileSize << 20
	c.Router = appRouter
}

func (c *Config) GetEnv(key, defaultValue string) string {
	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}

//pak evan aca
//02131999100
