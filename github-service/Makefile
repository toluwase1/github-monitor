run:
	go run cmd/github-monitor/main.go

# Makefile

# Variables
DOCKER_COMPOSE=docker-compose
ENV_FILE=.env

# Default target
.PHONY: all
all: up

# Build all services
.PHONY: build
build:
	$(DOCKER_COMPOSE) build

# Start all services
.PHONY: up
up: build
	$(DOCKER_COMPOSE) --env-file $(ENV_FILE) up -d

# Stop all services
.PHONY: down
down:
	$(DOCKER_COMPOSE) down

# View logs for all services
.PHONY: logs
logs:
	$(DOCKER_COMPOSE) logs -f

# Restart all services
.PHONY: restart
restart: down up

# Remove all Docker images and containers
.PHONY: clean
clean: down
	$(DOCKER_COMPOSE) rm -f
	docker system prune -f

# Check environment file
.PHONY: check-env
check-env:
ifndef ENV_FILE
	$(error ENV_FILE is undefined. Set ENV_FILE variable to your environment file path.)
endif
