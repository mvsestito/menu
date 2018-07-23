package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mvsestito/menu-api/api/storage"
)

// getItemsHandler handles GET requests to the /restaurant/{restaurantId}/item endpoint
func getItemsHandler(w http.ResponseWriter, req *http.Request) {
	var (
		err          error
		body         []byte
		items        []storage.Item
		restaurantID int
	)

	// parse uri vars
	vars := mux.Vars(req)
	if vars["restaurantId"] == "" {
		http.Error(w, "invalid url. restaurantId cannot be empty", http.StatusBadRequest)
		return
	} else {
		restaurantID, err = strconv.Atoi(vars["restaurantId"])
		if err != nil || restaurantID == 0 {
			http.Error(w, "must specify a nonzero integer for field 'restaurantId'", http.StatusBadRequest)
			return
		}
	}

	// get data
	items, err = storage.GetAllItems(restaurantID, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	body, err = json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// success
	SendJSONResp(w, body, http.StatusOK)
}
