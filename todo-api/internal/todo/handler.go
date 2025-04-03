package todo

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler handles HTTP requests for todos
type Handler struct {
	service *Service
}

// NewHandler creates a new todo handler
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// NewRouter creates a new router for the todo API
func NewRouter(handler *Handler) *mux.Router {
	r := mux.NewRouter()
	
	// Add CORS middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			next.ServeHTTP(w, r)
		})
	})
	
	r.HandleFunc("/todos", handler.GetAllTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", handler.GetTodoByID).Methods("GET")
	r.HandleFunc("/todos", handler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handler.UpdateTodo).Methods("PATCH")
	r.HandleFunc("/todos/{id}", handler.DeleteTodo).Methods("DELETE")
	r.HandleFunc("/todos/{id}", handler.Options).Methods("OPTIONS")
	return r
}

// GetAllTodos handles GET /todos
func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.GetAllTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// GetTodoByID handles GET /todos/{id}
func (h *Handler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todo, err := h.service.GetTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// CreateTodo handles POST /todos
func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todoCreate TodoCreate
	if err := json.NewDecoder(r.Body).Decode(&todoCreate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.service.CreateTodo(todoCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo handles PATCH /todos/{id}
func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var todoUpdate TodoUpdate
	if err := json.NewDecoder(r.Body).Decode(&todoUpdate); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.service.UpdateTodo(id, todoUpdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodo handles DELETE /todos/{id}
func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.service.DeleteTodo(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Options handles OPTIONS requests
func (h *Handler) Options(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// SetupRoutes sets up the routes for the todo API
func (h *Handler) SetupRoutes(router *mux.Router) {
	// Add CORS middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers
			w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			
			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			
			// Call the next handler
			next.ServeHTTP(w, r)
		})
	})

	// ... your existing route setup
	router.HandleFunc("/todos", h.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos", h.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", h.GetTodoByID).Methods("GET")
	router.HandleFunc("/todos/{id}", h.UpdateTodo).Methods("PATCH")
	router.HandleFunc("/todos/{id}", h.DeleteTodo).Methods("DELETE")
} 