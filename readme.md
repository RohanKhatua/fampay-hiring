# Youtube API - Fampay Hiring

## Tech Stack

- Golang : Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.

- PostgreSQL : PostgreSQL is a powerful, open source object-relational database system known for reliability, feature robustness, and performance (hosted on Vercel).

- GORM : An ORM library for Golang.

- Fiber : Fiber is an Express.js inspired web framework written in Go

## Setup

1. Clone the repository

```bash
$ git clone https://github.com/RohanKhatua/fampay-hiring

$ cd fampay-hiring
```

2. Install dependencies

```bash
$ go mod download
```

3. Create an env file and add the following variables

```bash
API_KEYS="<key1>,<key2>,..."
POSTGRES_USER="<USER>"
POSTGRES_HOST="<HOST>"
POSTGRES_PASSWORD="<PASSWORD>"
POSTGRES_DATABASE="<DATABASE NAME>"
POSTGRES_PORT="5432"
MIGRATE="true"
```

4. Run the server

```bash
$ go run main.go
```

## API Endpoints

1. To get the list of videos in reverse chronological order of their publishing date:

```bash
GET /api/videos
```

This defaults to 10 videos per page and returns the first page of the results.

To get the next page of results, use the `page` query parameter. To change the page size, use the `pageSize` query parameter.

```bash
GET /api/videos?page=2&pageSize=20
```

2. To search for videos by their title, description or channel:

```bash
GET /api/videos/search?q=<search_query>
```

Made with ❤️ and Go by Rohan Khatua
