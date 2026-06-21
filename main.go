package main

import (
	"fmt"
	"log"

	"github.com/alirfanyasin/golang-starterkit/config"
	"github.com/alirfanyasin/golang-starterkit/database"
	"github.com/alirfanyasin/golang-starterkit/packages/post"
	"github.com/alirfanyasin/golang-starterkit/routes"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	// 2. Connect to Database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// 3. Run Auto Migrations
	log.Println("Running database migrations...")
	err = db.AutoMigrate(&post.Post{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Database migrations completed successfully.")

	// 4. Setup Routes & Run Server
	r := routes.SetupRouter(db, cfg)
	port := fmt.Sprintf(":%s", cfg.AppPort)
	log.Printf("Server is running on port %s", cfg.AppPort)
	if err := r.Run(port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}