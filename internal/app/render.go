package app

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, any interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(any)
}

func RenderJSONwithStatus(w http.ResponseWriter, status int, any interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(any)
}

func RenderErrorJSON(w http.ResponseWriter, err *Error) {
	w.Header().Set("Content-Type", "application/json")
	if err.message == "" {
		err.message = err.err
	}
	if err.code == 0 {
		err.code = http.StatusBadRequest
	}
	w.WriteHeader(err.code)
	json.NewEncoder(w).Encode(*err)
}
