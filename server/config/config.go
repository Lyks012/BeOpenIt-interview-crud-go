package config

import (
	"database/sql"
)

func DBConnect() (db *sql.DB) {
	db, err := sql.Open("mysql", "user:user@tcp(db:3306)/crud-go")
	if err != nil {
		panic(err.Error())
	}

	return db
}
