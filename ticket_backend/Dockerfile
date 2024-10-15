# Use Go version 1.23
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /src/app

# Install air for live-reloading
RUN go install github.com/air-verse/air@latest

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./

# Download all Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Ensure all necessary Go files are present and tidy up the modules
RUN go mod tidy

# Build the Go application (assuming your main.go is located at cmd/api/main.go)
# Modify the path if your main.go is in a different location
RUN go build -o main cmd/api/main.go

# Set air to run as the entrypoint
CMD ["air"]
