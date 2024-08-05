# commit-fetcher/Dockerfile

# Use the official Golang image as the base image
FROM golang:1.19

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Install dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o commit-fetcher ./cmd/commit-fetcher

# Expose the port the app runs on
EXPOSE 8080

# Run the executable
CMD ["./commit-fetcher"]
