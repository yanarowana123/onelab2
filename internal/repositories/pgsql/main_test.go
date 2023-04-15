package pgsql

import (
	"database/sql"
	"github.com/subosito/gotenv"
	"github.com/yanarowana123/onelab2/configs"
	"log"
	"path/filepath"
	"testing"
)

var db *sql.DB

func TestMain(m *testing.M) {
	dir, err := filepath.Abs(filepath.Join(filepath.Dir("./"), "../../../"))

	err = gotenv.Load(filepath.Join(dir, ".env.test"))
	if err != nil {
		log.Fatal(err)
	}
	config, err := configs.New()

	if err != nil {
		log.Fatal(err)
	}
	db = ConnectDB(config.PgSqlDSN)

	m.Run()
}
