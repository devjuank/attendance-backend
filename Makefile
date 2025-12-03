.PHONY: help run build test test-coverage migrate-up migrate-down swagger lint docker-build docker-up docker-down clean

# Variables
BINARY_NAME=server
MAIN_PATH=cmd/server/main.go

help: ## Mostrar ayuda
	@echo "Comandos disponibles:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Ejecutar servidor en desarrollo
	@echo "🚀 Starting server..."
	go run $(MAIN_PATH)

build: ## Compilar binario
	@echo "🔨 Building binary..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)
	@echo "✅ Binary created at bin/$(BINARY_NAME)"

test: ## Ejecutar tests
	@echo "🧪 Running tests..."
	go test -v ./...

test-coverage: ## Tests con coverage
	@echo "🧪 Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "✅ Coverage report generated: coverage.html"

migrate-up: ## Ejecutar migraciones
	@echo "⬆️  Running migrations..."
	go run migrations/migrate.go up

migrate-down: ## Revertir migraciones
	@echo "⬇️  Reverting migrations..."
	@echo "TODO: Implement migrations rollback"

swagger: ## Generar documentación Swagger
	@echo "📖 Generating Swagger docs..."
	swag init -g $(MAIN_PATH) -o ./docs
	@echo "✅ Swagger docs generated"

lint: ## Ejecutar linter
	@echo "🔍 Running linter..."
	golangci-lint run

docker-build: ## Construir imagen Docker
	@echo "🐳 Building Docker image..."
	docker build -t attendance-backend:latest .

docker-up: ## Levantar Docker Compose
	@echo "🐳 Starting Docker Compose..."
	docker-compose up -d
	@echo "✅ Services started"

docker-down: ## Detener Docker Compose
	@echo "🐳 Stopping Docker Compose..."
	docker-compose down

clean: ## Limpiar archivos generados
	@echo "🧹 Cleaning..."
	rm -rf bin/
	rm -f coverage.out coverage.html
	rm -rf docs/swagger.json docs/swagger.yaml
	@echo "✅ Cleaned"

deps: ## Instalar dependencias
	@echo "📦 Installing dependencies..."
	go mod download
	go mod tidy
	@echo "✅ Dependencies installed"

dev: ## Ejecutar con hot reload (requiere air)
	@echo "🔥 Starting with hot reload..."
	air

install-tools: ## Instalar herramientas de desarrollo
	@echo "🛠️  Installing development tools..."
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "✅ Tools installed"
