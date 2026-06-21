package seeders

import (
	"log"

	"gorm.io/gorm"
)

// Seeder represents a database seeding function
type Seeder func(db *gorm.DB) error

// SeedAll runs all registered seeders
func SeedAll(db *gorm.DB) error {
	log.Println("Seeding database...")

	// List all seeders to be executed
	allSeeders := []Seeder{
		SeedPosts,
		// SeedUsers, // Add future seeders here
	}

	for _, seed := range allSeeders {
		if err := seed(db); err != nil {
			return err
		}
	}

	log.Println("Database seeding completed.")
	return nil
}
