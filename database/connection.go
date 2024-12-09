package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib" // PostgreSQL driver for Go
)

var DB *sql.DB

func InitDB() {
	var err error
	dbURL := os.Getenv("DB_URL") // Fetch DB URL from the environment
	if dbURL == "" {
		log.Fatal("DB_URL is not set in environment variables")
	}

	DB, err = sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connection established")
}
