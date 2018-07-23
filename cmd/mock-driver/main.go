package main

import (
	"log"

	"github.com/mvsestito/menu-api/api/storage"
	"github.com/mvsestito/menu-api/api/storage/mock"
)

func main() {
	mock.ResetTables(storage.DB)
	mock.AddMockRestaurants(storage.DB)
	mock.AddMockItems(storage.DB)
	log.Println("mock data successfully added to db")
}
