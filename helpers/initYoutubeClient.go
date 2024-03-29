package helpers

import (
	"log"
	"net/http"
	"os"
	"strings"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var currentKeyIndex int = 0

func getNextApiKey(apiKeys []string) string {
	// Cycle through API keys when the current one is exhausted
	if currentKeyIndex >= len(apiKeys) {
		currentKeyIndex = 0
	}

	key := apiKeys[currentKeyIndex]
	currentKeyIndex++
	return key
}

func InitYoutubeClient() *youtube.Service {
	// Gets the next API key from the list of API keys
	apiKey := getNextApiKey(strings.Split(os.Getenv("API_KEYS"), ","))

	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)

	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	return service
}
