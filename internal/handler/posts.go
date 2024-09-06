package handler

import (
	"net/http"
	"strconv"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
	"github.com/PeterKWIlliams/feed-aggregator-go/internal/models"
)

func (c *Config) HandlerGetPosts(w http.ResponseWriter, r *http.Request, user database.User) {
	limitString := r.URL.Query().Get("limit")
	limit := 10
	if limitString != "" {
		newLimit, err := strconv.Atoi(limitString)
		limit = newLimit
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Invalid limit parameter default used")
		}
	}
	databasePosts, err := c.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Couldnt get posts")
		return
	}
	posts := models.DatabasePostsToPosts(databasePosts)
	RespondWithJSON(w, http.StatusOK, posts)
}
