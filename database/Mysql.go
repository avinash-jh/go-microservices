package database

import (
	"fmt"
	"product-api/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectAndMigrate() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Retrieve underlying sql.DB instance for pool configuration
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	// Configure connection pool settings
	sqlDB.SetMaxOpenConns(10)                  // Maximum open connections
	sqlDB.SetMaxIdleConns(5)                   // Keep 5 idle connections warm
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // Close connection after 30 mins
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // Close connection if idle for 10 mins

	// Verify connection
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}
	fmt.Println("Database connection established")

	// AutoMigrate creates the table if it does not exist
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w", err)
	}
	fmt.Println("Product table created successfully")

	return db, nil
}
