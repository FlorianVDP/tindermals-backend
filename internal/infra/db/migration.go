package db

import (
	"gorm.io/gorm"
	userinfra "jamlink-backend/internal/modules/user/infra"
	"log"
)

func MigrateDB(db *gorm.DB) {
	log.Println("🚀 Running global database migrations...")

	userinfra.MigrateUserTable(db)

	log.Println("✅ All migrations completed successfully!")
}
