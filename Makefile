#локально
run:
	go run cmd/main.go -local
#докер
upbuild:
	docker-compose up --build avito-backend-test
up:
	docker-compose up avito-backend-test
build:
	docker build -t avito-backend-test .
prune:
	docker container prune
dockerrun:
	docker run --name=avito-backend-test -p 8090:8090 --rm avito-backend-test
down:
	docker-compose down avito-backend-test

.PHONY: run push prune build dockerrun upbuild up down
