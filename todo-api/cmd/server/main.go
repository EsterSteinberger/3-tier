package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"todo-api/internal/database"
	"todo-api/internal/todo"
)

func main() {
	// Get environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}
	
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "todouser"
	}
	
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "todopass"
	}
	
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "tododb"
	}
	
	// Initialize the database
	db, err := database.NewPostgresDB(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	
	// Initialize the todo service with the database
	todoService := todo.NewService(db)
	
	// Initialize the todo handler
	todoHandler := todo.NewHandler(todoService)
	
	// Create a new router
	router := mux.NewRouter()
	
	// Set up the routes
	todoHandler.SetupRoutes(router)
	
	// Start the server
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 