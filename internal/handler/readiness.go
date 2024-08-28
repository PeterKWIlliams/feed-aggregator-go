package handler

import "net/http"

func (cfg *Config) HandlerHealth(w http.ResponseWriter, r *http.Request) {
	RespondWithJSON(w, http.StatusOK, "SERVER HEALTH OK")
}
