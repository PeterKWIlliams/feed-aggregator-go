package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
)

type Feed struct {
	ID            uuid.UUID  `json:"id"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Name          string     `json:"name"`
	Url           string     `json:"url"`
	UserID        uuid.UUID  `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func DatabaseFeedtoFeed(feed database.Feed) Feed {
	var lastFetchedAt *time.Time
	if feed.LastFetchedAt.Valid {
		lastFetchedAt = &feed.LastFetchedAt.Time
	}

	return Feed{
		ID:            feed.ID,
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
		Name:          feed.Name,
		Url:           feed.Url,
		UserID:        feed.UserID,
		LastFetchedAt: lastFetchedAt,
	}
}

func DatabaseFeedsToFeeds(databaseFeeds []database.Feed) []Feed {
	var feeds []Feed
	for _, feed := range databaseFeeds {
		feeds = append(feeds, DatabaseFeedtoFeed(feed))
	}
	fmt.Println(len(feeds))
	return feeds
}
