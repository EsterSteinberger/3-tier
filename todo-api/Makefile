.PHONY: build run test clean docker-build docker-run

# Go related variables
BINARY_NAME=todo-api
MAIN_FILE=cmd/server/main.go

# Build the application
build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

# Run the application
run: build
	./$(BINARY_NAME)

# Run tests
test:
	go test -v ./...

# Clean up
clean:
	go clean
	rm -f $(BINARY_NAME)

# Build Docker image
docker-build:
	docker build -t todo-api .

# Run Docker container
docker-run: docker-build
	docker run -p 8080:8080 todo-api