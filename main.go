package main

import (
	"log"
	"yt_api/database"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// apiKey := os.Getenv("API_KEY")

	// client := &http.Client{
	// 	Transport: &transport.APIKey{Key: apiKey},
	// }

	// service, err := youtube.New(client)

	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	database.ConnectToDB()

	// arr := []string{"id", "snippet"}

	// call := service.Search.List(arr).Q("Google").MaxResults(25).Type("video")

	// response, err := call.Do()

	// if err != nil {
	// 	log.Fatalf("Error making search API call: %v", err)
	// }

	// for _, item := range response.Items {
	// 	log.Printf("Video: %v, %v", item.Snippet.Title, item.Id.VideoId)
	// }
}
