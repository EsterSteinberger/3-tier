FROM golang:1.20-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-api ./cmd/server

# Use a minimal alpine image
FROM alpine:3.17

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /todo-api /todo-api

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["/todo-api"] 