package handler

import "net/http"

func (cfg *Config) HandlerHealth(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Status string `json:"status"`
	}
	RespondWithJSON(w, http.StatusOK, response{Status: "OK"})
}
