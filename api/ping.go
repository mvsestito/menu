package api

import (
	"fmt"
	"net/http"

	"github.com/mvsestito/menu-api/api/storage"
)

// pingHandler handles requests to /ping for server health checks
func pingHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
}

// pingDbHandler handles requests to /ping/db for database health checks
func pingDbHandler(w http.ResponseWriter, req *http.Request) {
	if err := storage.DB.Ping(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		SendTextResp(w, "OK", http.StatusOK)
	}
}
