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
			call := service.Search.List(callArr).Q("asmr").MaxResults(10).Type("video")
			response, err := call.Do()

			if err != nil {
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

				videosBatch = append(videosBatch, *video)

				if len(videosBatch) >= batchSize {
					StoreVideos(videosBatch)
					videosBatch = []models.Video{} // Reset the batch
				}
			}

			// Store any remaining videos in the last batch
			if len(videosBatch) > 0 {
				StoreVideos(videosBatch)
			}
		}
	}
}
