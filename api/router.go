package api

import (
	"github.com/gorilla/mux"
)

var ROUTER *mux.Router

func initRouter() {
	ROUTER = mux.NewRouter()

	// ping routes
	ROUTER.HandleFunc("/ping", pingHandler).Name("ping")
	ROUTER.HandleFunc("/ping/", pingHandler).Name("ping")
	ROUTER.HandleFunc("/ping/db", pingDbHandler).Name("ping-db")
	ROUTER.HandleFunc("/ping/db/", pingDbHandler).Name("ping-db")

	// GET requests
	ROUTER.HandleFunc("/restaurant/{restaurantId}/item", getItemsHandler).
		Methods("GET").Name("get-items")
}
