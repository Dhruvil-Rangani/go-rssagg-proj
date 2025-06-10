# Go RSS Aggregator

A backend web service written in Go that aggregates RSS feeds for users. Users can register, subscribe to feeds, and retrieve personalized aggregated posts. The service continuously fetches the latest articles from subscribed feeds in the background and stores them in a PostgreSQL database.

## 🚀 Features

- User registration with API key-based authentication.
- Add new RSS feeds and follow/unfollow them.
- Periodic background scraping of RSS feeds using Go’s concurrency features (goroutines and WaitGroups).
- RESTful API built with the Chi router, using JSON responses.
- PostgreSQL storage with proper relational schema and migrations via Goose.
- Type-safe SQL queries generated with SQLC for clean database interactions.

## 🛠️ Tech Stack

- **Go** (1.19+)
- **Chi** (HTTP router and middleware)
- **PostgreSQL** (database)
- **Goose** (database migrations)
- **SQLC** (type-safe query generation)
- **UUID** (unique IDs for all entities)
- **Goroutines** (for concurrent feed scraping)
- **CORS** (Cross-Origin Resource Sharing)

## 📦 Setup Instructions

1. **Clone the repository**:

    ```bash
    git clone https://github.com/Dhruvil-Rangani/go-rssagg-proj.git
    cd go-rssagg-proj
    ```

2. **Set environment variables**:

    Create a `.env` file in the project root:

    ```env
    PORT=8080
    DB_URL=postgres://username:password@localhost:5432/rssagg?sslmode=disable
    ```

3. **Install dependencies**:

    Ensure Go modules are up-to-date:

    ```bash
    go mod tidy
    ```

4. **Run database migrations**:

    Make sure Goose CLI is installed (`go install github.com/pressly/goose/v3/cmd/goose@latest`).

    ```bash
    goose -dir sql/schema postgres "$DB_URL" up
    ```

5. **Run the server**:

    ```bash
    go run main.go
    ```

## 🗂️ Database Schema

- **users**: stores user data, including unique API key for authentication.
- **feeds**: stores RSS feed metadata.
- **feeds_follows**: maps users to the feeds they follow (many-to-many).
- **posts**: stores articles fetched from feeds.

## 🌐 API Usage

### Authentication

Use the `Authorization` header with the format:
```
Authorization: `ApiKey <api_key>`
```

### Endpoints

- **Health Check**:  
    `GET /v1/healthz`

- **Register User**:  
    `POST /v1/users`

    ```json
    {
      "name": "username"
    }
    ```

- **Add Feed**:  
    `POST /v1/feeds`

    ```json
    {
      "name": "Feed Name",
      "url": "http://example.com/rss"
    }
    ```

- **Get All Feeds**:  
    `GET /v1/feeds`

- **Follow Feed**:  
    `POST /v1/feeds/follows`

    ```json
    {
      "feed_id": "<feed_uuid>"
    }
    ```

- **Unfollow Feed**:  
    `DELETE /v1/feeds/follows/{feedFollowID}`

- **Get User’s Posts**:  
    `GET /v1/users/posts`

## 🚀 Deployment

- Configure `PORT` and `DB_URL` environment variables.
- Compile the server:  
    ```bash
    go build -o rssagg .
    ./rssagg
    ```
- Run behind a reverse proxy for HTTPS (recommended).
- Use process managers like systemd or Docker for production deployment.

## 🤝 Contributing

Contributions welcome! Please fork the repo, make changes, and open a pull request. For larger changes, please open an issue first to discuss your ideas.

---

Built by Dhruvil Rangani.
