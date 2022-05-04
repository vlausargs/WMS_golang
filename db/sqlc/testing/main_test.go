package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	sqlc "github.com/hidratechh/GO_WMS/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:6543/noqueue?sslmode=disable"
)

var testQuery *sqlc.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQuery = sqlc.New(conn)

	os.Exit(m.Run())
}
