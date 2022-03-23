package datastore

import (
	"log"

	"gorm.io/gorm"

	"blog/models"
)

func autoMigrate(orm *gorm.DB) {
	log.Println("Auto-migration start...")

	// Try auto-migrate database tables with specified entities.
	if err := orm.AutoMigrate(models.Note{}); err != nil {
		log.Fatal("automigrate failed:", err)
	}

	log.Println("Auto-migration complete.")
}
