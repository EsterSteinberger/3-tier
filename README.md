# TODO Application

A modern, containerized TODO application with a Go backend API, PostgreSQL database, and React frontend.

## Features

- Create, read, update, and delete TODO items
- Mark TODOs as completed
- Persistent storage with PostgreSQL
- Clean, modern UI with responsive design
- Containerized with Docker for easy deployment

## Tech Stack

- **Backend**: Go with Gorilla Mux for routing
- **Database**: PostgreSQL
- **Frontend**: React with modern CSS
- **Infrastructure**: Docker, Docker Compose, Nginx

## Project Structure 

```
.
├── docker-compose.yml          # Docker Compose configuration
├── todo-api/                   # Go backend
│   ├── cmd/server/main.go      # Main application file
│   ├── internal/               # Internal packages
│   │   ├── database/           # Database implementations
│   │   └── todo/               # Todo domain logic
│   ├── go.mod                  # Go module definition
│   └── go.sum                  # Go module checksums
└── todo-frontend/              # React frontend
    ├── public/                 # Static assets
    ├── src/                    # Source code
    │   ├── components/         # React components
    │   ├── services/           # API services
    │   ├── styles/             # CSS styles
    │   └── config.js           # Configuration
    ├── Dockerfile              # Docker configuration
    └── nginx.conf              # Nginx configuration
```

## Getting Started

### Prerequisites

- Docker and Docker Compose

### Running the Application

1. Clone the repository
2. Start the application:
   ```bash
   docker-compose up -d
   ```
3. Access the application at http://localhost

## API Configuration

The frontend communicates with the backend API through Nginx proxy:

```javascript
export const API_CONFIG = {
  BASE_URL: '/api',
};
```

## Database

The application uses PostgreSQL for data persistence with todos table storing:
- ID
- Title
- Completed status
- Created and updated timestamps

## Docker Configuration

The application is containerized using Docker with the following services:

- **postgres**: PostgreSQL database
- **todo-api**: Go backend API
- **todo-frontend**: React frontend with Nginx

## Future Enhancements

- User authentication
- Todo categories or tags
- Due dates for todos
- Search and filtering capabilities

---

Created with Docker, Go, React, and PostgreSQL




