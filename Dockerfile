# Build stage
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-s' -o /home/billing-engine ./cmd/main.go

# Final stage
FROM alpine:latest

# Install necessary packages
RUN apk add curl

# Create a new directory to store the application
RUN mkdir -p /home/config

WORKDIR /home

# Copy the built binary and other necessary files
COPY --from=builder /home/billing-engine ./
COPY internal/config/*.json ./config

# Expose the application port
EXPOSE 8000

# Run the application with nodemon
CMD ["/home/billing-engine"]