package todo

import (
	"errors"
)

// Repository defines the interface for todo storage
type Repository interface {
	GetAllTodos() ([]Todo, error)
	GetTodoByID(id string) (Todo, error)
	CreateTodo(todo TodoCreate) (Todo, error)
	UpdateTodo(id string, todo TodoUpdate) (Todo, error)
	DeleteTodo(id string) error
}

// Service handles the business logic for todos
type Service struct {
	repo Repository
}

// NewService creates a new todo service
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetAllTodos returns all todos
func (s *Service) GetAllTodos() ([]Todo, error) {
	return s.repo.GetAllTodos()
}

// GetTodoByID returns a todo by its ID
func (s *Service) GetTodoByID(id string) (Todo, error) {
	if id == "" {
		return Todo{}, errors.New("id cannot be empty")
	}
	return s.repo.GetTodoByID(id)
}

// CreateTodo creates a new todo
func (s *Service) CreateTodo(todoCreate TodoCreate) (Todo, error) {
	if todoCreate.Title == "" {
		return Todo{}, errors.New("title cannot be empty")
	}
	return s.repo.CreateTodo(todoCreate)
}

// UpdateTodo updates an existing todo
func (s *Service) UpdateTodo(id string, todoUpdate TodoUpdate) (Todo, error) {
	if id == "" {
		return Todo{}, errors.New("id cannot be empty")
	}
	return s.repo.UpdateTodo(id, todoUpdate)
}

// DeleteTodo deletes a todo by its ID
func (s *Service) DeleteTodo(id string) error {
	if id == "" {
		return errors.New("id cannot be empty")
	}
	return s.repo.DeleteTodo(id)
} 