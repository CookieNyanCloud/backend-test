run:
	go run cmd/main.go -local

upbuild:
	docker-compose up --build backend-test
up:
	docker-compose up backend-test
build:
	docker build -t backend-test .
prune:
	docker container prune
dockerrun:
	docker run --name=backend-test -p 8090:8090 --rm backend-test
down:
	docker-compose down backend-test

.PHONY: run push prune build dockerrun upbuild up down
