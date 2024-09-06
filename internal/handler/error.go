package handler

import "net/http"

func (cfg *Config) HandlerError(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
