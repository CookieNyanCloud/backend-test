
run:
	go run cmd/main.go -local

push:
	git push origin master

prune:
	docker container prune

build:
	docker build -t avito-backend-test .

dockerrun:
	docker run --name=avito-backend -p 8090:8090 --rm avito-backend

upbuild:
	docker-compose up --build avito-backend
up:
	docker-compose up avito-backend
down:
	docker-compose down

.PHONY: run push prune build dockerrun upbuild up down
