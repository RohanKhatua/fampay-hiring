package helpers

import (
	"log"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func InitYoutubeClient() *youtube.Service {
	apiKey := os.Getenv("API_KEY")

	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)

	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	return service
}
