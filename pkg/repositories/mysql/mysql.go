package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/yanarowana123/onelab2/configs"
)

// зачем это здесь ?
// pkg = это общая папка которую другие могут импортировать к себе.
func New(config configs.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.SqlDSN)
	return db, err
}
