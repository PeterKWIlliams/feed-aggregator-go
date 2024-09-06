package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
)

type Post struct {
	ID          uuid.UUID  `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Title       string     `json:"title"`
	Url         string     `json:"url"`
	Description *string    `json:"description"`
	PublishedAt *time.Time `json:"published_at"`
	FeedID      uuid.UUID  `json:"feed_id"`
}

func DatabasePostToPost(databasePost database.Post) Post {
	return Post{
		ID:          databasePost.ID,
		CreatedAt:   databasePost.CreatedAt,
		UpdatedAt:   databasePost.UpdatedAt,
		Title:       databasePost.Title,
		Url:         databasePost.Url,
		Description: StrPtrFromNullString(databasePost.Description),
		PublishedAt: TimePtrFromNullTime(databasePost.PublishedAt),
		FeedID:      databasePost.FeedID,
	}
}

func DatabasePostsToPosts(databasePosts []database.Post) []Post {
	var posts []Post
	for _, databasePost := range databasePosts {
		posts = append(posts, DatabasePostToPost(databasePost))
	}
	return posts
}

func StrPtrFromNullString(nullString sql.NullString) *string {
	if nullString.Valid {
		return &nullString.String
	}
	return nil
}

func TimePtrFromNullTime(nullTime sql.NullTime) *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}
	return nil
}
