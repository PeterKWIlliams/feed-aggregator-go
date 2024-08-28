package handler

import (
	"encoding/json"
	"net/http"
)

type Config struct{}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(code)
}

func ResponsWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	type ErrorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, ErrorResponse{Error: message})
}
