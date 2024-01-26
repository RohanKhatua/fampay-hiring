package helpers

import (
	"fmt"
	"time"
	"yt_api/models"

	"google.golang.org/api/youtube/v3"
)

func ParseVideo(snippet youtube.SearchResultSnippet, videoId string) (*models.Video, error) {
	timeLayout := "2006-01-02T15:04:05Z"
	parsedTime, err := time.Parse(timeLayout, snippet.PublishedAt)
	if err != nil {
		return nil, fmt.Errorf("error parsing time: %v", err)
	}

	return &models.Video{
		ID:               videoId,
		VideoTitle:       snippet.Title,
		VideoDescription: snippet.Description,
		ChannelTitle:     snippet.ChannelTitle,
		PublishTime:      parsedTime,
	}, nil
}
