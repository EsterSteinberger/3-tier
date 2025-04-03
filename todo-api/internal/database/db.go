package database

import (
	"errors"
	"sync"
	"time"
	"todo-api/internal/todo"

	"github.com/google/uuid"
)

// InMemoryDB is a simple in-memory database implementation
type InMemoryDB struct {
	todos map[string]todo.Todo
	mu    sync.RWMutex
}

// NewInMemoryDB creates a new in-memory database
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		todos: make(map[string]todo.Todo),
	}
}

// GetAllTodos returns all todos from the database
func (db *InMemoryDB) GetAllTodos() ([]todo.Todo, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	todos := make([]todo.Todo, 0, len(db.todos))
	for _, t := range db.todos {
		todos = append(todos, t)
	}
	return todos, nil
}

// GetTodoByID returns a todo by its ID
func (db *InMemoryDB) GetTodoByID(id string) (todo.Todo, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	t, exists := db.todos[id]
	if !exists {
		return todo.Todo{}, errors.New("todo not found")
	}
	return t, nil
}

// CreateTodo creates a new todo
func (db *InMemoryDB) CreateTodo(todoCreate todo.TodoCreate) (todo.Todo, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	now := time.Now()
	newTodo := todo.Todo{
		ID:        uuid.New().String(),
		Title:     todoCreate.Title,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	db.todos[newTodo.ID] = newTodo
	return newTodo, nil
}

// UpdateTodo updates an existing todo
func (db *InMemoryDB) UpdateTodo(id string, todoUpdate todo.TodoUpdate) (todo.Todo, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	t, exists := db.todos[id]
	if !exists {
		return todo.Todo{}, errors.New("todo not found")
	}

	if todoUpdate.Title != nil {
		t.Title = *todoUpdate.Title
	}
	if todoUpdate.Completed != nil {
		t.Completed = *todoUpdate.Completed
	}
	t.UpdatedAt = time.Now()

	db.todos[id] = t
	return t, nil
}

// DeleteTodo deletes a todo by its ID
func (db *InMemoryDB) DeleteTodo(id string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.todos[id]; !exists {
		return errors.New("todo not found")
	}

	delete(db.todos, id)
	return nil
} 