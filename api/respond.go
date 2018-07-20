package api

import (
	"net/http"
)

// SendTextResp sends an http response to client with plaintext body
func SendTextResp(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(code)

	if msg != "" {
		bytes := []byte(msg)
		w.Write(bytes)
	}
}

// SendJSONResp sends an http response to client with a JSON encoded body
func SendJSONResp(w http.ResponseWriter, data []byte, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
