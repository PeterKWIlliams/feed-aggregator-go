package models

import (
	"time"

	"github.com/google/uuid"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FeedID    uuid.UUID `json:"feed_id"`
	UserID    uuid.UUID `json:"user_id"`
}

func DatabaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
		FeedID:    feedFollow.FeedID,
		UserID:    feedFollow.UserID,
	}
}

func DatabaseFeedFollowsToFeedFollows(databaseFeedFollows []database.FeedFollow) []FeedFollow {
	var feedFollows []FeedFollow
	for _, feedFollow := range databaseFeedFollows {
		feedFollows = append(feedFollows, DatabaseFeedFollowToFeedFollow(feedFollow))
	}
	return feedFollows
}
