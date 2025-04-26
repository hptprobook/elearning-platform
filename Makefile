.PHONY: build test clean docker-build docker-push

# Build tất cả services
build:
	@echo "Building all services..."
	@cd cmd/auth && go build
	@cd cmd/content && go build
	@cd cmd/user && go build

# Test tất cả services
test:
	@echo "Testing all services..."
	@cd cmd/auth && go test ./...
	@cd cmd/content && go test ./...
	@cd cmd/user && go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf cmd/*/bin

# Build Docker images
docker-build:
	@echo "Building Docker images..."
	docker-compose build

# Push Docker images
docker-push:
	@echo "Pushing Docker images..."
	docker-compose push

# Run services locally
run:
	@echo "Running services..."
	docker-compose up

# Stop services
stop:
	@echo "Stopping services..."
	docker-compose down