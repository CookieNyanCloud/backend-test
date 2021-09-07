
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

#createdb:
#	docker exec -it avito-backend-test_db_1 createdb --username=postgres postgres

#migrateup:
#	#migrate -path avito-backend-test/schema -database "postgres" -verbose up
#	docker run -v {{ migration dir }}:/migrations --network host migrate/migrate
#    -path=/migrations/ -database postgres://localhost:5432/database up 2
#
#
#migratedown:
#	migrate -path avito-backend-test/schema -database "postgres" -verbose down


.PHONY: run push prune build dockerrun upbuild up down
