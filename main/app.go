package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	guuid "github.com/google/uuid"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"root:P@ssw0rd@tcp(127.0.0.1:3306)/enigma")
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)
	if err != nil {
		log.Fatal(err)
	}

	// Simple query
	rows, err := db.Query("select id,first_name,last_name,address,city from m_customer where first_name like ?", "Ka")
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(rows)

	var id, firstName, lastName, address, city string
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &address, &city)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, firstName, lastName, address, city)
	}

	// Single Row
	var totalRecord int
	err = db.QueryRow("select count(*) from m_customer").Scan(&totalRecord)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(totalRecord)

	//Insert, Delete, Update
	//newCustomerId := guuid.New()
	//_, err = db.Exec("insert into m_customer values (?,?,?,?,?)", newCustomerId, "Maysista", "Deviani", "Ciracas", "Jakarta")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Insert Success")

	//Transactional
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer func(tx *sql.Tx) {
		if err := recover(); err != nil {
			err := tx.Rollback()
			if err != nil {
				log.Fatal(err)
			}
		}
	}(tx)
	newCustomerId := guuid.New()
	_, err = tx.Exec("insert into m_customer values (?,?,?,?,?)", newCustomerId, "Tika", "Yesi", "Ragunan", "Jakarta")

	newCustomerId = guuid.New()
	_, err = tx.Exec("insert into m_customer values (?,?,?,?,?)", newCustomerId, "Jution", "Chandra", "Ragunan", "Jakarta")

	//Simulate error->rollback
	//panic("Failed connection")
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}
