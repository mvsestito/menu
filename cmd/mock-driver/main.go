package main

import (
	"database/sql"
	"log"

	"github.com/burntsushi/toml"
	"github.com/mvsestito/menu-api/api"
	"github.com/mvsestito/menu-api/api/storage/mock"
)

func main() {
	c := api.Config{}

	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal("error decoding TOML config: ", err)
	}

	db, err := sql.Open("postgres", c.Dbstr())
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	mock.ResetTables(db)
	mock.AddMockRestaurants(db)
	mock.AddMockItems(db)

	log.Println("mock data successfully added to db")
}
