package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/PeterKWIlliams/feed-aggregator-go/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading.env file: %v", err)
	}
	srv := server.NewServer()
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
