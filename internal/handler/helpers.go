package handler

import (
	"encoding/json"
	"net/http"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
)

type Config struct {
	DB *database.Queries
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(code)
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	type ErrorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJSON(w, code, ErrorResponse{Error: message})
}
