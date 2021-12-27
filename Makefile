run:
	go run cmd/main.go -local
up:
	docker-compose up --build backend-test
docker-run:
	docker run --name=backend-test -p 8090:8090 --rm backend-test
down:
	docker-compose down backend-test

mock:
	go generate -v ./...

.PHONY: run docker-run up down mock
