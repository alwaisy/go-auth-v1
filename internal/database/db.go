package database

import (
	"fmt"
	"go-auth-v1/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB() {
	cfg := config.LoadConfig(".")

	dsn := cfg.Database.DatabaseUrl

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")
}

// CloseDB closes the database connection gracefully
func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB() // Get underlying *sql.DB
		if err != nil {
			log.Printf("Error getting database instance: %v", err)
			return
		}

		fmt.Println("Closing database connection...")
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			fmt.Println("Database connection closed successfully.")
		}
	}
}
