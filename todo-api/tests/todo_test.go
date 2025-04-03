package tests

import (
	"testing"
	"todo-api/internal/database"
	"todo-api/internal/todo"
)

func TestCreateTodo(t *testing.T) {
	// Initialize the database
	db := database.NewInMemoryDB()

	// Initialize the service
	service := todo.NewService(db)

	// Create a todo
	todoCreate := todo.TodoCreate{
		Title: "Test Todo",
	}

	createdTodo, err := service.CreateTodo(todoCreate)
	if err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	if createdTodo.Title != todoCreate.Title {
		t.Errorf("Expected title %s, got %s", todoCreate.Title, createdTodo.Title)
	}

	if createdTodo.Completed {
		t.Errorf("Expected completed to be false, got true")
	}
}

func TestGetTodoByID(t *testing.T) {
	// Initialize the database
	db := database.NewInMemoryDB()

	// Initialize the service
	service := todo.NewService(db)

	// Create a todo
	todoCreate := todo.TodoCreate{
		Title: "Test Todo",
	}

	createdTodo, err := service.CreateTodo(todoCreate)
	if err != nil {
		t.Fatalf("Failed to create todo: %v", err)
	}

	// Get the todo by ID
	retrievedTodo, err := service.GetTodoByID(createdTodo.ID)
	if err != nil {
		t.Fatalf("Failed to get todo by ID: %v", err)
	}

	if retrievedTodo.ID != createdTodo.ID {
		t.Errorf("Expected ID %s, got %s", createdTodo.ID, retrievedTodo.ID)
	}

	if retrievedTodo.Title != createdTodo.Title {
		t.Errorf("Expected title %s, got %s", createdTodo.Title, retrievedTodo.Title)
	}
} 