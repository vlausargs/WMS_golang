package main

import (
	"database/sql"
	"log"

	api_rest "github.com/hidratechh/GO_WMS/api/rest"
	db "github.com/hidratechh/GO_WMS/db/sqlc/exec"

	_ "github.com/lib/pq"
)

const (
	dbDriver1      = "postgres"
	dbSource1      = "postgresql://root:secret@localhost:6543/noqueue?sslmode=disable"
	serverAddress1 = "0.0.0.0:8080"
)

func main() {

	/* db and api setup*/
	store := startDatabase1()
	startRestAPI(store)

}

func startDatabase1() *db.SQLStore {
	conn, err := sql.Open(dbDriver1, dbSource1)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	return store
}

func startRestAPI(store *db.SQLStore) {
	server := api_rest.NewServer(store)

	err := server.Start(serverAddress1)
	if err != nil {
		log.Fatal("cannot start the server: ", err)
	}
}
