package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	pingHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, "OK", string(body))
}

func TestPingDbHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	pingDbHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, "OK", string(body))
}
