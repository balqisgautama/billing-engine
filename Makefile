# Variables  
APP_NAME = billing-engine
DOCKER_IMAGE = $(APP_NAME):latest  
DOCKERFILE = Dockerfile  
CONTAINER_NAME = billing-engine
  
# Default target  
all: build  
  
# Build the Docker image  
build:  
	docker build -t $(DOCKER_IMAGE) -f $(DOCKERFILE) .  
  
# Run the Docker container  
run: 
	nodemon
  
# Stop and remove the Docker container  
stop:  
	docker-compose down  
  
# Clean up Docker images and containers  
clean: stop
	docker rm -f $$(docker ps -aq) || true  
	docker rmi -f $(DOCKER_IMAGE) || true  
	docker rmi -f billing-engine-billing-engine-app || true  

rebuild: stop clean run

# Run tests  
# test:  
# 	go test ./... -coverprofile=coverage.out
#     go tool cover -func=coverage.out
#     go tool cover -html=coverage.out -o coverage.html

# test-coverage: test
# 	go tool cover -func=coverage.out | grep total | awk '{print $$3}' > coverage.txt

# Migration
# db-migrate:
# 	go run cmd/db_migrate/main.go