package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mvsestito/menu-api/api/storage"
)

// GetItemsHandler handles GET requests to the /restaurant/{restaurantId}/item endpoint
func GetItemsHandler(w http.ResponseWriter, req *http.Request) {
	var (
		err   error
		body  []byte
		items []storage.Item
	)

	// parse uri vars
	vars := mux.Vars(req)
	restaurantId = vars["resturantId"]
	if restaurantId == 0 {
		http.Error(w, "invalid url. restaurantId must be greater than zero", http.StatusBadRequest)
		return
	}

	// get data
	items, err = storage.GetAllItems(Db, restaurantId, "")
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
