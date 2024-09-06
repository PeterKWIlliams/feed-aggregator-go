package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
	"github.com/PeterKWIlliams/feed-aggregator-go/internal/models"
)

func (c *Config) HandlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Couldn't decode request body")
	}
	feedFollow, err := c.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Feed follow creation failed")
		return
	}
	RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

func (c *Config) HandlerGetUsersFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := c.DB.GetFeedFollowsByUserId(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid feedFollowID")
	}
	RespondWithJSON(w, http.StatusOK, models.DatabaseFeedFollowsToFeedFollows(feedFollows))
}

func (c *Config) HandlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID, err := uuid.Parse(r.PathValue("feedFollowID"))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}
	_, err = c.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Feed follow deletion failed")
		return
	}
	RespondWithJSON(w, http.StatusNoContent, nil)
}
