postgres:
		docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
		docker exec -it postgres12 createdb --username=root --owner=root fishing_with

dropdb:
		docker exec -it postgres12 dropdb -U root fishing_with

migrateup:
		@.	./app.env;	\
		migrate -path db/migrate -database "$${DB_SOURCE}" -verbose up

migratedown:
		@.	./app.env;	\
		migrate -path db/migrate -database "$${DB_SOURCE}" -verbose down

migrateup1:
		@.	./app.env;	\
		migrate -path db/migrate -database "$${DB_SOURCE}" -verbose up 1
		
migratedown1:
		@.	./app.env;	\
		migrate -path db/migrate -database "$${DB_SOURCE}" -verbose down 1

sqlc:
		sqlc generate

test:
		go test -v -cover ./...

server:
		go run main.go

mock:
		mockgen -package mockdb -destination db/mock/store.go github.com/tomoropy/fishing-with-backend/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1 

