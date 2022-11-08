package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tomoropy/fishing-with-backend/api"
	db "github.com/tomoropy/fishing-with-backend/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/fishing_with?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot conenct to db:", err)
	}

	store := db.New(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalln("cannot start server:", err)
	}
}
