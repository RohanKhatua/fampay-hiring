package models

import (
	"time"
)

type Video struct {
	ID               string `json:"id" gorm:"primarykey"`
	CreatedAt        time.Time
	VideoTitle       string    `json:"video_title"`
	VideoDescription string    `json:"video_description"`
	ChannelTitle     string    `json:"channel_title"`
	PublishTime      time.Time `json:"publish_time" gorm:"index:idx_publish_time;index:idx_publish_time_id,priority:2"`
}
