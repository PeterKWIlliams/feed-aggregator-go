package server

import (
	"net/http"
	"os"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/handler"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()
	port := ":" + os.Getenv("PORT")
	cfg := &handler.Config{}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	mux.HandleFunc("/healthz", cfg.HandlerHealth)
	return &http.Server{
		Addr:    port,
		Handler: mux,
	}
}
