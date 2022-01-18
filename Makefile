run:
	go run cmd/main.go -local
up:
	docker-compose up --build backend-test
build:
	docker-compose up backend-test
mock:
	go generate -v ./...
test:
	go test ./... -v
lint:
	golangci-lint run

.PHONY: run up build down mock test lint
