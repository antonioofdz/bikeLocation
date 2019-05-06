package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", "dev-user:dev-password@/ProyectoFinalDra")
	if err != nil {
		return db, err
	}

	return db, nil
}