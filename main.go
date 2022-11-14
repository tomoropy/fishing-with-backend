package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/tomoropy/fishing-with-backend/api"
	db "github.com/tomoropy/fishing-with-backend/db/sqlc"
	"github.com/tomoropy/fishing-with-backend/util"
)


func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannnot load confg:", err)
	}
	
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalln("cannot conenct to db:", err)
	}

	store := db.New(conn)
	server, err := api.NewServer(config, store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("cannot start server:", err)
	}
}
