package database

import (
	"fmt"
	"log"
	"os"
	"yt_api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

func GlobalActivationScope(db *gorm.DB) *gorm.DB {
	return db.Where("is_activated = ?", true)
}

var Database DbInstance

func ConnectToDB() {

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false,
	})

	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	migrateFlag := os.Getenv("MIGRATE")

	if migrateFlag == "true" {
		db.AutoMigrate(&models.Video{})
	}

	Database = DbInstance{Db: db}
}
