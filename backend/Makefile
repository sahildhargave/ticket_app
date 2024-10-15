# Start the application with Docker Compose
start:
	@docker-compose up --build

# Stop the application and remove the containers
stop:
	@docker-compose down -v --remove-orphans
	@docker rmi ticket-booking

# Build the Go application
build:
	@go build -o ./bin/main ./cmd/api/main.go
