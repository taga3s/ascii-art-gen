PROJECT_NAME=ascii-art-gen

build:
		@echo "Building the project..."
		@go build -o bin/$(PROJECT_NAME) ./main.go

test:
		@echo "Running tests..."
		@go test -v ./...