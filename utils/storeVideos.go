package utils

import (
	"log"
	"yt_api/database"
	"yt_api/models"

	"gorm.io/gorm/clause"
)

func StoreVideos(videos []models.Video) {
	if len(videos) == 0 {
		return
	}

	// Insert all videos in a single transaction
	result := database.Database.Db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&videos)
	if result.Error != nil {
		log.Printf("Error storing videos: %v", result.Error)
	}
}
