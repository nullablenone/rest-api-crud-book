package main

import (
	"log"
	"net/http"

	"rest-api-crud-books/config"
	"rest-api-crud-books/models"
	"rest-api-crud-books/routes"
)

func main() {
	// Initialize database
	db := config.ConnectDB()
	defer func() {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}()

	// Run migrations
	models.Migrate(db)
	log.Println("Database migration completed successfully!")

	// Setup routes
	router := routes.SetupRoutes(db)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
