package utils

import (
	"log"
	"time"
	"yt_api/helpers"
	"yt_api/models"

	"google.golang.org/api/youtube/v3"
)

func FetchVideos(service *youtube.Service) {
	ticker := time.NewTicker(10 * time.Second)
	callArr := []string{"id", "snippet"}
	batchSize := 10

	for {
		select {
		case <-ticker.C:
			call := service.Search.List(callArr).Q("asmr").MaxResults(10).Type("video").Order("date").PublishedAfter(time.Now().AddDate(0, 0, -5).Format(time.RFC3339))
			response, err := call.Do()

			// Handle API quota exceeded error
			if err.Error() == "googleapi: Error 403: The request cannot be completed because you have exceeded your <a href=\"/youtube/v3/getting-started#quota\">quota</a>., quotaExceeded" {
				log.Printf("Quota exceeded, switching to next API key")
				service = helpers.InitYoutubeClient()
				continue
			} else if err != nil {
				log.Printf("Error making search API call: %v", err)
				continue
			}

			var videosBatch []models.Video

			for _, item := range response.Items {
				video, err := helpers.ParseVideo(*item.Snippet, item.Id.VideoId)
				if err != nil {
					log.Printf("%v", err)
					continue
				}

				// Add video to batch
				videosBatch = append(videosBatch, *video)

				if len(videosBatch) >= batchSize {
					go StoreVideos(videosBatch)    // Using a goroutine
					videosBatch = []models.Video{} // Reset the batch
				}
			}

			// Store any remaining videos in the last batch
			if len(videosBatch) > 0 {
				go StoreVideos(videosBatch) // Using a goroutine
			}
		}
	}
}
