# feed-aggregator-go

A web server that allows clients to add, follow, and fetch RSS feeds.

## Features

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds added by other users
- Fetch the latest posts from the RSS feeds you follow

## Getting Started

### Installation

1. Clone the repository:

   ```
   git clone github.com/PeterKWIlliams/feed-aggregator-go
   ```

2. Navigate to the project directory:

   ```
   cd feed-aggregator-go
   ```

3. Create a .env file in the root directory of the project:

   ```
    touch .env
   ```

   Add the following environment variables:

   - `PORT`="" # Specify the port number for the server to listen on
   - `DB_URL`="" # Provide the connection string for your postgres database

4. Build the project:

   ```
   go build -o out ./cmd/server/main.go
   ```

5. Run the server:

   ```
   ./out
   ```

   The server will start running on `http://localhost:8080` (or configured port in env)

## Usage

Authorization is provided by sending along the users api key in the `Authorization` header.
The api key is generated automatically on user creation.

Once the server is running, you can interact with it using the following endpoints:

- **`GET /healthz`**: Health check.
- **`POST /v1/users`**: Add a new user.
- **`GET /v1/users`**: Get user details (authenticated).
- **`POST /v1/feeds`**: Create a new feed (authenticated).
- **`GET /v1/feeds`**: Get all feeds.
- **`POST /v1/feed_follows`**: Follow a feed (authenticated).
- **`GET /v1/feed_follows`**: Get user's feed follows (authenticated).
- **`DELETE /v1/feed_follows/{feedFollowID}`**: Unfollow a feed (authenticated).
- **`GET /v1/posts`**: Get latest posts from followed feeds (authenticated).
