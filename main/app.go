package main

import (
	"enigmacamp.com/gosql/config"
	"enigmacamp.com/gosql/repositories"
	"enigmacamp.com/gosql/usecases"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
)

/*
1. viper config reader => go get github.com/spf13/viper
2. kingpin cli parser => go get gopkg.in/alecthomas/kingpin.v2
3. Buat folder config, config.go
4. Buat folder repositories, lalu db_session_factory.go
5. Buat folder usecase
*/

var (
	appEnvironment = kingpin.Flag("env", "Environment").HintOptions("dev", "prod").Default("dev").Short('e').String()
)

type app struct {
	sf              *repositories.DbSessionFactory
	customerUseCase usecases.ICustomerUseCase
}

func newApp() app {
	kingpin.Parse()
	c := config.NewConfig(*appEnvironment)
	err := c.InitDb()
	if err != nil {
		panic(err)
	}
	myapp := app{
		sf:              c.SessionFactory,
		customerUseCase: usecases.NewCustomerUseCase(c.SessionFactory),
	}
	return myapp
}
func (a app) run() {
	totalCustomer, err := a.customerUseCase.GetTotalCustomer()
	log.Printf("Total Customer : %v \n", totalCustomer)
	customers, err := a.customerUseCase.FindCustomerByFirstName("ka")
	log.Printf("Customer : %v \n", customers)

	if err != nil {
		log.Println(err)
	}
}
func main() {
	newApp().run()
}

//func main() {
//	db, err := sql.Open("mysql",
//		"root:P@ssw0rd@tcp(127.0.0.1:3306)/enigma")
//	defer func(db *sql.DB) {
//		err := db.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(db)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Simple query
//	rows, err := db.Query("select id,first_name,last_name,address,city from m_customer where first_name like ?", "Ka")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer func(rows *sql.Rows) {
//		err := rows.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(rows)
//
//	var id, firstName, lastName, address, city string
//	for rows.Next() {
//		err := rows.Scan(&id, &firstName, &lastName, &address, &city)
//		if err != nil {
//			log.Fatal(err)
//		}
//		log.Println(id, firstName, lastName, address, city)
//	}
//
//	// Single Row
//	var totalRecord int
//	err = db.QueryRow("select count(*) from m_customer").Scan(&totalRecord)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println(totalRecord)
//
//	//Insert, Delete, Update
//	//newCustomerId := guuid.New()
//	//_, err = db.Exec("insert into m_customer values (?,?,?,?,?)", newCustomerId, "Maysista", "Deviani", "Ciracas", "Jakarta")
//	//if err != nil {
//	//	log.Fatal(err)
//	//}
//	//log.Println("Insert Success")
//
//	//Transactional
//	tx, err := db.Begin()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer func(tx *sql.Tx) {
//		if err := recover(); err != nil {
//			err := tx.Rollback()
//			if err != nil {
//				log.Fatal(err)
//			}
//		}
//	}(tx)
//	newCustomerId := guuid.New()
//	_, err = tx.Exec("insert into m_customer values (?,?,?,?,?)", newCustomerId, "Tika", "Yesi", "Ragunan", "Jakarta")
//
//	newCustomerId = guuid.New()
//	_, err = tx.Exec("insert into m_customer values (?,?,?,?,?)", newCustomerId, "Jution", "Chandra", "Ragunan", "Jakarta")
//
//	//Simulate error->rollback
//	//panic("Failed connection")
//	err = tx.Commit()
//	if err != nil {
//		log.Fatal(err)
//	}
//}
