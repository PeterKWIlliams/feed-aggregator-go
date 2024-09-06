package server

import (
	"database/sql"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/database"
	"github.com/PeterKWIlliams/feed-aggregator-go/internal/handler"
	"github.com/PeterKWIlliams/feed-aggregator-go/internal/scraper"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()
	port := ":" + os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	dbQueries := database.New(db)
	apiCfg := &handler.Config{DB: dbQueries}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello, world!"))
	})
	mux.HandleFunc("/healthz", apiCfg.HandlerHealth)
	mux.HandleFunc("POST /v1/users", apiCfg.CreateUser)
	mux.HandleFunc("GET /v1/users", apiCfg.MiddlewareAuth(apiCfg.GetUser))
	mux.HandleFunc("POST /v1/feeds", apiCfg.MiddlewareAuth(apiCfg.CreateFeed))
	mux.HandleFunc("GET /v1/feeds", apiCfg.GetFeeds)
	mux.HandleFunc("POST /v1/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	mux.HandleFunc("GET /v1/feed_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUsersFeedFollow))
	mux.HandleFunc("GET /v1/posts", apiCfg.MiddlewareAuth(apiCfg.HandlerGetPosts))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollow))

	waitTime := time.Duration(120) * time.Second
	go scraper.StartScraping(dbQueries, 3, waitTime)

	return &http.Server{
		Addr:    port,
		Handler: mux,
	}
}
