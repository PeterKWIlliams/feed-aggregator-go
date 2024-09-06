package scraper

import (
	"context"
	"database/sql"
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
)

func StartScraping(db *database.Queries, fetchAmount int, waitTime time.Duration) {
	ticker := time.NewTicker(waitTime)
	for ; true; <-ticker.C {
		feedsTofetch, err := db.GetNextFeedsToFetch(context.Background(), int32(fetchAmount))
		if err != nil {
			log.Printf("Could not get feeds to fetch: %v", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feedsTofetch {
			wg.Add(1)
			go scrape(db, wg, feed)

		}
		wg.Wait()
	}
}

var Formats = [...]string{
	"Mon, 02 Jan 2006 15:04:05 -0700",
	"02 Jan 2006 15:04:05 -0700",
	"2006-01-02T15:04:05Z07:00",
	"Mon, 02 Jan 2006 15:04:05 MST",
	"02 Jan 2006",
}

func scrape(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error while updating feed to fetched: %v", err)
		return
	}
	data, err := fetchFeed(feed.Url)
	if err != nil {
		log.Printf("error while fetching feed: %v", err)
		return
	}

	for _, item := range data.Channel.Items {
		var parsedTime time.Time
		var err error
		for _, format := range Formats {
			parsedTime, err = time.Parse(format, item.PubDate)
			if err == nil {
				break
			}
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Description: sql.NullString{String: item.Description, Valid: true},
			Url:         item.Link,
			PublishedAt: sql.NullTime{Time: parsedTime, Valid: true},
			FeedID:      feed.ID,
		})
		if err != nil {
			log.Printf("error while creating post: %v Description length:%v", err, len(item.Description))
		}
	}
	log.Printf("Successfuly fetched: %v feeds", len(data.Channel.Items))
}

type RssFeed struct {
	Channel struct {
		Title       string        `xml:"title"`
		Link        string        `xml:"link"`
		Description string        `xml:"description"`
		Language    string        `xml:"language"`
		Items       []RssFeedItem `xml:"item"`
	} `xml:"channel"`
}

type RssFeedItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(url string) (*RssFeed, error) {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RssFeed
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}
