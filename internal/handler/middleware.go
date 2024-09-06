package handler

import (
	"net/http"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/auth"
	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (c *Config) MiddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken, err := auth.ExtractApiKey(r.Header)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Unauthorized invalid api key")
			return
		}
		user, err := c.DB.GetUserByApiKey(r.Context(), authToken)
		if err != nil {
			RespondWithError(w, http.StatusUnauthorized, "Couldn't find user with that api key")
			return
		}
		handler(w, r, user)
	}
}
