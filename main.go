package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"github.com/trenchesdeveloper/simplebank/api"
	db "github.com/trenchesdeveloper/simplebank/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	ServerAddress = "0.0.0.0:8080"
)

func main(){
	var err error
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}