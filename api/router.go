package api

import (
	"github.com/gorilla/mux"
)

var router *mux.Router

func initRouter() {
	router = mux.NewRouter()

	// ping routes
	router.HandleFunc("/ping", pingHandler).Name("ping")
	router.HandleFunc("/ping/", pingHandler).Name("ping")
	router.HandleFunc("/ping/db", pingOCDbHandler).Name("ping-db")
	router.HandleFunc("/ping/db/", pingOCDbHandler).Name("ping-db")

	// GET requests
	router.Handle("/restaurant/{restaurantId}/item", getItemsHandler).
		Methods("GET").Name("get-items")
}
