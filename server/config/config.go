package config

import("database/sql")

func DBConnect() (db *sql.DB){
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud-go")
	if err != nil{
		panic(err.Error())
	}

	return db
}

