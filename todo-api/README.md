# TODO API

A Go-based REST API for managing TODO items with PostgreSQL persistence.

## Quick Start with Makefile

```bash
# Run the API locally
make run

# Build the API
make build

# Run tests
make test

# Clean build artifacts
make clean

# Format code
make fmt

# Lint code
make lint

# Generate API documentation
make docs
```

## API Endpoints

| Method | Endpoint      | Description           |
|--------|---------------|-----------------------|
| GET    | /todos        | Get all todos         |
| GET    | /todos/{id}   | Get a specific todo   |
| POST   | /todos        | Create a new todo     |
| PATCH  | /todos/{id}   | Update a todo         |
| DELETE | /todos/{id}   | Delete a todo         |

## Environment Variables

- `PORT`: API server port (default: 8080)
- `DB_HOST`: PostgreSQL host
- `DB_PORT`: PostgreSQL port
- `DB_USER`: PostgreSQL username
- `DB_PASSWORD`: PostgreSQL password
- `DB_NAME`: PostgreSQL database name

## Development

```bash
# Install dependencies
go mod download

# Run with hot reload
make dev

# Run specific tests
go test ./internal/todo
```

## Docker

```bash
# Build Docker image
docker build -t todo-api .

# Run in Docker
docker run -p 8080:8080 todo-api
```

---

For more details, see the main [project README](../README.md). 