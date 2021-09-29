package repositories

import "database/sql"

type DbSessionFactory struct {
	Db *sql.DB
}

func NewDbSessionFactory(driverName, dataSourceName string) *DbSessionFactory {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	factory := new(DbSessionFactory)
	factory.Db = db
	return factory
}
