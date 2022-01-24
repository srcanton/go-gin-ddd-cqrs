CONNECTION=postgres:postgres@127.0.0.1:5432

build:
	docker-compose -f docker-compose.yml build

run: build
	docker-compose -f docker-compose.yml up -d

stop:
	docker-compose -f docker-compose.yml down

localrun:
	docker-compose -f docker-compose.local.yml up -d
	go run main.go

localstop:
	docker-compose -f docker-compose.local.yml down

createdb:
	docker exec -it gogindddcqrs-postgres createdb --username=postgres --owner=postgres gogindddcqrs

dropdb:
	docker exec -it gogindddcqrs-postgres dropdb gogindddcqrs

migrateup:
	migrate -path src/infrastructure/common/persistence/migration -database 'postgresql://$(CONNECTION)/gogindddcqrs?sslmode=disable' -verbose up

migratedown:
	migrate -path src/infrastructure/common/persistence/migration -database 'postgresql://$(CONNECTION)/gogindddcqrs?sslmode=disable' -verbose down