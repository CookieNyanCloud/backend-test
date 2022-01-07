run:
	go run cmd/main.go -local
up:
	docker-compose up --build --force-recreate backend-test

build:
	docker-compose up backend-test

mock:
	go generate -v ./...

.PHONY: run up build down mock
