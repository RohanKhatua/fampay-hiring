package routes

import (
	"yt_api/database"
	"yt_api/models"

	"github.com/gofiber/fiber/v2"
)

func SearchVideos(c *fiber.Ctx) error {
	query := c.Query("q")

	var videos []models.Video
	result := database.Database.Db.
		Where("video_title ILIKE ?", "%"+query+"%").
		Or("video_description ILIKE ?", "%"+query+"%").
		Or("channel_title ILIKE ?", "%"+query+"%").
		Find(&videos)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot search videos",
		})
	}

	return c.JSON(videos)
}
