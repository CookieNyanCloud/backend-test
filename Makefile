
run:
	go run cmd/main.go -local

push:
	git push origin master

prune:
	docker container prune

build:
	docker build -t avito-backend-test .

dockerrun:
	docker run --name=avito-backend-test -p 8090:8090 --rm avito-backend-test

upbuild:
	docker-compose up --build avito-backend-test
up:
	docker-compose up avito-backend-test
down:
	docker-compose down

.PHONY: run push prune build dockerrun upbuild up down
