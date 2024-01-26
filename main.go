package main

import (
	"log"
	"yt_api/database"
	"yt_api/helpers"
	"yt_api/routes"
	"yt_api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/videos", routes.GetVideos)
	app.Get("/api/videos/search", routes.SearchVideos)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectToDB()

	app := fiber.New()
	setupRoutes(app)

	// Initialize YouTube client
	youtubeClient := helpers.InitYoutubeClient()

	// Run FetchVideos in a separate goroutine
	go utils.FetchVideos(youtubeClient)

	// Start Fiber server
	log.Fatal(app.Listen(":3000"))
}
