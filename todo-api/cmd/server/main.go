package main

import (
	"log"
	"net/http"
	"os"
	"todo-api/internal/database"
	"todo-api/internal/todo"
)

func main() {
	// Initialize the database
	db := database.NewInMemoryDB()

	// Initialize the service
	service := todo.NewService(db)

	// Initialize the handler
	handler := todo.NewHandler(service)

	// Initialize the router
	router := todo.NewRouter(handler)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
} 