package pgsql

import (
	"database/sql"
	"github.com/lib/pq"
)

func ConnectDB(pgDSN string) *sql.DB {
	pgUrl, err := pq.ParseURL(pgDSN)

	if err != nil {
		panic(err)
	}

	db, err := sql.Open("postgres", pgUrl)

	if err != nil {
		panic(err)
	}

	return db
}
