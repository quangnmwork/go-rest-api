package utils

import (
	"fmt"
	"todo-app/models"
)

func AutoMigrate() error {
	// Connect to database
	db, err := ConnectDB()
	if err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
		return err
	}

	fmt.Println("Connected to database successfully")
	fmt.Println("Starting database migrations...")

	// Migrate User model first
	if err := db.AutoMigrate(&models.User{}, &models.Token{}); err != nil {
		fmt.Printf("Failed to migrate User model: %v\n", err)
		return err
	}
	fmt.Println("User model migrated successfully")

	// Check if Token model exists before migrating it
	// Comment out the Token migration if it doesn't exist yet
	/*
		if err := db.AutoMigrate(&models.Token{}); err != nil {
			fmt.Printf("Failed to migrate Token model: %v\n", err)
			return err
		}
		fmt.Println("Token model migrated successfully")
	*/

	fmt.Println("All migrations completed successfully")
	return nil
}
