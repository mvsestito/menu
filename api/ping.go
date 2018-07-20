package api

import (
	"fmt"
	"net/http"
)

// pingHandler handles requests to /ping for server health checks
func pingHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "OK")
}

// pingDbHandler handles requests to /ping/db for database health checks
func pingDbHandler(w http.ResponseWriter, req *http.Request) {
	var (
		ok  string
		err error
	)

	err = DB.QueryRow(`SELECT 'OK'`).Scan(&ok)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if ok == "" {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	} else {
		SendTextResp(w, "OK", http.StatusOK)
		return
	}
}
