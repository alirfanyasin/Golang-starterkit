package migrations

import (
	"log"

	"gorm.io/gorm"
)

// MigrateAll runs auto-migrations for all registered models
func MigrateAll(db *gorm.DB) error {
	log.Println("Running database migrations...")

	var models []interface{}

	// Register models from separate migration files
	models = append(models, GetPostModels()...)
	// models = append(models, GetUserModels()...) // Add future models here

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	log.Println("Database migrations completed successfully.")
	return nil
}
