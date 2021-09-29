package config

import (
	repo "enigmacamp.com/gosql/repositories"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type dbConf struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	schema     string
	dbEngine   string
}

type Config struct {
	SessionFactory *repo.DbSessionFactory
	dbConf         *dbConf
}

func NewConfig(env string) *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName(env)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Can not read config")
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
	return c
}

func (c *Config) InitDb() error {
	log.Println("======= Create DB Connection =======")
	sf := repo.NewDbSessionFactory(c.dbConf.dbEngine, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.dbConf.dbUser, c.dbConf.dbPassword, c.dbConf.dbHost, c.dbConf.dbPort, c.dbConf.schema))
	c.SessionFactory = sf
	return nil
}

func (c *Config) GetEnv(key, defaultValue string) string {
	if envVal := viper.GetString(key); len(envVal) != 0 {
		return envVal
	}
	return defaultValue
}
