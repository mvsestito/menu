package api

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func initDB() {
	var (
		err error
		row string
	)

	log.Println("Connecting to DB...")
	Db, err = sql.Open("postgres", *flagDBStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %s", err)
	}

	err = Db.QueryRow("SELECT 'OK'").Scan(&row)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}

	Db.SetMaxIdleConns(5)   // limit idle conns
	Db.SetMaxOpenConns(100) // limit open conns

	log.Println("Connected to DB")
}
