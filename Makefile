run:
	go run cmd/main.go -local
up:
	docker-compose up --build backend-test
build:
	docker-compose up backend-test
prune:
	docker image prune
mock:
	go generate -v ./...
test:
	go test ./... -v

.PHONY: run up build down mock test prune
