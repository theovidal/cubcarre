package db

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(os.Getenv("DB_PATH")), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v\n", err)
	}
	db.AutoMigrate()
	return db
}
