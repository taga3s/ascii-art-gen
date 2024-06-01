PROJECT_NAME=ascii-art-gen

build:
		@echo "Building the project..."
		@go build -o bin/$(PROJECT_NAME) ./main.go
		@echo "Build successful!"

test:
		@echo "Running tests..."
		@go test -v ./...

release-check:
		goreleaser check

release-snapshot:
		goreleaser release --snapshot --clean