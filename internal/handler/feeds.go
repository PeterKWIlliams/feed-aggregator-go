package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
	"github.com/PeterKWIlliams/feed-aggregator-go/internal/models"
)

func (c *Config) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	type response struct {
		Feed       models.Feed       `json:"feed"`
		FeedFollow models.FeedFollow `json:"feed_follow"`
	}
	params := parameters{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Couldn't decode request body")
		return
	}
	databaseFeed, err := c.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Couldnt create feed")
		return
	}
	databaseFeedFollow, err := c.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    databaseFeed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Feed follow creation failed")
	}
	feed := models.DatabaseFeedtoFeed(databaseFeed)
	feedFollow := models.DatabaseFeedFollowToFeedFollow(databaseFeedFollow)

	RespondWithJSON(w, http.StatusCreated, response{Feed: feed, FeedFollow: feedFollow})
}

func (c *Config) GetFeeds(w http.ResponseWriter, r *http.Request) {
	databaseFeeds, err := c.DB.GetFeeds(r.Context())
	feeds := models.DatabaseFeedsToFeeds(databaseFeeds)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Feeds retrieval failed")
		return
	}
	RespondWithJSON(w, http.StatusAccepted, feeds)
}
