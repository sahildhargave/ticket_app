package db

import (
	"fmt"
	"log" // Use the standard log package

	"github.com/sahildhargave/ticket-project-v1/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Init initializes the database connection
func Init(config *config.EnvConfig, DBMigrator func(*gorm.DB) error) (*gorm.DB, error) {
	uri := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=%s port=5432",
		config.DBHost, config.DBUser, config.DBName, config.DBPassword, config.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err // Return error
	}

	log.Println("Connected to the database")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate: %v", err)
		return nil, err // Return error
	}

	return db, nil // Return db and nil for error
}
