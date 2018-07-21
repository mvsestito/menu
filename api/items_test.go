package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mvsestito/menu-api/api/storage"
	"github.com/mvsestito/menu-api/api/storage/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetItemsHandler(t *testing.T) {
	assert := assert.New(t)
	DB = mock.MockDB()
	mock.ResetTables(DB)
	mock.AddMockRestaurants(DB)
	mock.AddMockItems(DB)

	// restaurant 1
	endpoint := "/restaurant/{restaurantId}/item"
	req := httptest.NewRequest("GET", endpoint, nil)
	req = mux.SetURLVars(req, map[string]string{"restaurantId": "1"})
	w := httptest.NewRecorder()
	getItemsHandler(w, req)

	resp := w.Result()

	var items []storage.Item
	err := json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		assert.Nil(err)
		return
	}

	assert.Equal(6, len(items))

	// restaurant 2
	req = httptest.NewRequest("GET", endpoint, nil)
	req = mux.SetURLVars(req, map[string]string{"restaurantId": "2"})
	w = httptest.NewRecorder()
	getItemsHandler(w, req)

	resp = w.Result()

	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		assert.Nil(err)
		return
	}

	assert.Equal(4, len(items))

	// empty restaurant id
	req = httptest.NewRequest("GET", endpoint, nil)
	req = mux.SetURLVars(req, map[string]string{"restaurantId": ""})
	w = httptest.NewRecorder()
	getItemsHandler(w, req)

	resp = w.Result()
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(http.StatusBadRequest, resp.StatusCode)
	assert.Equal("invalid url. restaurantId cannot be empty\n", string(body))
}
