package main

import (
	"log"
	"yt_api/database"
	"yt_api/helpers"
	"yt_api/utils"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	database.ConnectToDB()

	utils.FetchVideos(helpers.InitYoutubeClient())
}
