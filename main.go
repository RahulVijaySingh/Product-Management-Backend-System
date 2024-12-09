package main

import (
	"log"
	"product-management-system/database"
	"product-management-system/router"
	"product-management-system/cache"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	database.InitDB()
	cache.InitRedis()

	// Start the router
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
