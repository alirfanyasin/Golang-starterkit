package main

import (
	"fmt"
	"log"

	"github.com/alirfanyasin/golang-starterkit/config"
	"github.com/alirfanyasin/golang-starterkit/database"
	"github.com/alirfanyasin/golang-starterkit/database/migrations"
	"github.com/alirfanyasin/golang-starterkit/database/seeders"
	"github.com/alirfanyasin/golang-starterkit/routes"
)

// @title           Golang Starterkit API
// @version         1.0
// @description     RESTful API starterkit using Gin and GORM.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Type "Bearer " followed by your JWT token.
func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Connect to Database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// 3. Run Migrations & Seeders
	if err := migrations.MigrateAll(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	if err := seeders.SeedAll(db); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}

	// 4. Setup Routes & Run Server
	r := routes.SetupRouter(db, cfg)
	port := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server is running on port %s", cfg.AppPort)
	if err := r.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}