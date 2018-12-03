package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func db_exec(db *sql.DB, q string) {
	var _, err = db.Exec(q)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	os.Create("./data.db")

	var db *sql.DB

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var q = ""

	q = "CREATE TABLE memo ("
	q += " id INTEGER PRIMARY KEY AUTOINCREMENT"
	q += ", body VARCHAR(255) NOT NULL"
	q += ", created_at TIMESTAMP DEFAULT (DATETIME('now','localtime'))"
	q += ")"
	db_exec(db, q)

	q = "INSERT INTO memo "
	q += " (body)"
	q += " VALUES"
	q += " ('body1')"
	q += ",('body2')"
	db_exec(db, q)

	db.Close()
}
