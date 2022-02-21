NAME=star-wars
OS ?= linux

.PHONY: compile
compile:
	@echo "=> installing dependencies..."
	go mod tidy
	@echo "==> Compiling star-wars..."
	go build -o build/server cmd/api/main.go

.PHONY: dev-local
dev-local:
	@echo ">>>>> Starting server application..."
	go mod tidy
	docker-compose up --d redis mongodb
	go run cmd/api/main.go

.PHONY: dev-docker-up
dev-docker-up:
	@echo ">>>>> Starting server application..."
	docker-compose up --build -d

.PHONY: dev-docker-down
dev-docker-down:
	@echo ">>>>> Shutting application..."
	docker-compose down

.PHONY: test
test:
	@echo "==> Running tests..."
	go test ./... --race -count=1 -v

.PHONY: test-coverage
test-coverage:
	@echo "==> Running tests..."
	go test -failfast --race -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

.PHONY: lint
lint:
	@echo "Running golangci-lint"
	@golangci-lint run --fix

.PHONY: generate
generate:
	@go mod tidy
	@go install github.com/matryer/moq@latest
	@echo "==>cleaning up generated files"
	find . -type f -name '*_mock.go' -exec rm {} +
	@echo "==>running go generate..."
	go generate ./...
	swag init -g ./cmd/api/main.go -o ./docs/swagger
