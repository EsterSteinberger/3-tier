package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"todo-api/internal/todo"
)

// PostgresDB implements the Repository interface for PostgreSQL
type PostgresDB struct {
	db *sql.DB
}

// NewPostgresDB creates a new PostgreSQL database connection
func NewPostgresDB(host, port, user, password, dbname string) (*PostgresDB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	
	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	// Create the todos table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			completed BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		)
	`)
	if err != nil {
		return nil, err
	}
	
	return &PostgresDB{db: db}, nil
}

// GetAllTodos retrieves all todos from the database
func (p *PostgresDB) GetAllTodos() ([]todo.Todo, error) {
	rows, err := p.db.Query(`
		SELECT id, title, completed, created_at, updated_at 
		FROM todos 
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var todos []todo.Todo
	for rows.Next() {
		var t todo.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	
	return todos, nil
}

// GetTodoByID retrieves a todo by its ID
func (p *PostgresDB) GetTodoByID(id string) (todo.Todo, error) {
	var t todo.Todo
	err := p.db.QueryRow(`
		SELECT id, title, completed, created_at, updated_at 
		FROM todos 
		WHERE id = $1
	`, id).Scan(&t.ID, &t.Title, &t.Completed, &t.CreatedAt, &t.UpdatedAt)
	
	if err != nil {
		return todo.Todo{}, err
	}
	
	return t, nil
}

// CreateTodo creates a new todo
func (p *PostgresDB) CreateTodo(todoCreate todo.TodoCreate) (todo.Todo, error) {
	newTodo := todo.Todo{
		ID:        uuid.New().String(),
		Title:     todoCreate.Title,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	_, err := p.db.Exec(`
		INSERT INTO todos (id, title, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
	`, newTodo.ID, newTodo.Title, newTodo.Completed, newTodo.CreatedAt, newTodo.UpdatedAt)
	
	if err != nil {
		return todo.Todo{}, err
	}
	
	return newTodo, nil
}

// UpdateTodo updates an existing todo
func (p *PostgresDB) UpdateTodo(id string, todoUpdate todo.TodoUpdate) (todo.Todo, error) {
	// First check if the todo exists
	existingTodo, err := p.GetTodoByID(id)
	if err != nil {
		return todo.Todo{}, err
	}
	
	// Update the todo
	if todoUpdate.Title != nil {
		existingTodo.Title = *todoUpdate.Title
	}
	if todoUpdate.Completed != nil {
		existingTodo.Completed = *todoUpdate.Completed
	}
	existingTodo.UpdatedAt = time.Now()
	
	// Save the updated todo
	_, err = p.db.Exec(`
		UPDATE todos
		SET title = $1, completed = $2, updated_at = $3
		WHERE id = $4
	`, existingTodo.Title, existingTodo.Completed, existingTodo.UpdatedAt, id)
	
	if err != nil {
		return todo.Todo{}, err
	}
	
	return existingTodo, nil
}

// DeleteTodo deletes a todo
func (p *PostgresDB) DeleteTodo(id string) error {
	result, err := p.db.Exec(`DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("todo with ID %s not found", id)
	}
	
	return nil
}

// Close closes the database connection
func (p *PostgresDB) Close() error {
	return p.db.Close()
} 