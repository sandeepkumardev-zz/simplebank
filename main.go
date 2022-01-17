package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sandeepkumardev/simplebank/api"
	db "github.com/sandeepkumardev/simplebank/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:pswd@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "localhost:3000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	if err := server.Start(serverAddress); err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
