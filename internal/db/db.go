package db

import (
	"fmt"
	"go-auth-v1/internal/config"
	"go-auth-v1/internal/domain/auth"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// ConnectDB initializes and stores the global DB instance
func ConnectDB() {
	cfg := config.LoadConfig(".")
	dsn := cfg.Database.DatabaseUrl

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	log.Println("Database connected successfully")
}

func InitDB() (*gorm.DB, error) {
	cfg := config.LoadConfig(".")
	dsn := cfg.Database.DatabaseUrl

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	// Run migrations
	err = db.AutoMigrate(&auth.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}

// CloseDB closes the db connection gracefully
func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB() // Get underlying *sql.DB
		if err != nil {
			log.Printf("Error getting db instance: %v", err)
			return
		}

		fmt.Println("Closing db connection...")
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing db connection: %v", err)
		} else {
			fmt.Println("Database connection closed successfully.")
		}
	}
}
