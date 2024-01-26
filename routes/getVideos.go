package routes

import (
	"strconv"
	"yt_api/database"
	"yt_api/models"

	"github.com/gofiber/fiber/v2"
)

func GetVideos(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "10"))

	var videos []models.Video

	// Pagination through Offset and Limit
	result := database.Database.Db.Order("publish_time desc").Offset((page - 1) * pageSize).Limit(pageSize).Find(&videos)

	if len(videos) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "No videos found",
		})
	}

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting videos",
		})
	}

	return c.JSON(videos)

}
