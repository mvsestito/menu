package api

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func initDB() {
	var (
		err error
		row string
	)

	log.Println("Connecting to DB...")
	DB, err = sql.Open("postgres", CONFIG.Dbstr())
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	err = DB.QueryRow("SELECT 'OK'").Scan(&row)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	log.Println("Connected to DB")
}
