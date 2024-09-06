# feed-aggregator-go

A web server that allows clients to add, follow, and fetch RSS feeds.

## Features

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds added by other users
- Fetch the latest posts from the RSS feeds you follow
- Keep up with your favorite blogs, news sites, podcasts, and more!

## Getting Started

### Prerequisites

- Go (version X.X or later)

### Installation

1. Clone the repository:

   ```
   "github.com/PeterKWIlliams/feed-aggregator-go"
   ```

2. Navigate to the project directory:

   ```
   cd feed-aggregator-go
   ```

3. Build the project:

   ```
   go build -o out ./cmd/server/main.go
   ```

4. Run the server:

   ```
   ./out
   ```

   The server will start running on `http://localhost:8080` (or configured port in env)

## Usage

Once the server is running, you can interact with it using the following endpoints:

- `POST /feeds`: Add a new RSS feed
- `GET /feeds`: Get a list of all available RSS feeds
- `POST /follows`: Follow an RSS feed
- `DELETE /follows`: Unfollow an RSS feed
- `GET /posts`: Get the latest posts from the RSS feeds you follow

