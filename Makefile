postgres:
		docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
		docker exec -it postgres12 createdb --username=root --owner=root fishing_with

dropdb:
		docker exec -it postgres12 dropdb --username=root --owner=root fishing_with

migrateup:
		migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/fishing_with?sslmode=disable" -verbose up

migratedown:
		migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/fishing_with?sslmode=disable" -verbose down

sqlc:
		sqlc generate

test:
		go test -v -cover ./...

server:
		go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server

