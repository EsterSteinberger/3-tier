package todo

import (
	"time"
)

// Todo represents a todo item in the application
type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TodoCreate is used for creating a new todo
type TodoCreate struct {
	Title string `json:"title"`
}

// TodoUpdate is used for updating an existing todo
type TodoUpdate struct {
	Title     *string `json:"title,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
} 