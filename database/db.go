package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// データベースのコネクションを開く
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// テーブル作成
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "BOOKS" ("ID" INTEGER PRIMARY KEY, "TITLE" VARCHAR(255))`,
	)
	if err != nil {
		panic(err)
	}

	// データの挿入
	res, err := db.Exec(
		`INSERT INTO BOOKS (ID, TITLE) VALUES (?, ?)`,
		123,
		"title",
	)
	if err != nil {
		panic(err)
	}
	// 挿入処理の結果からIDを取得
	id, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("inesrt id>%d", id)

	// 複数レコード取得
	rows, err := db.Query(
		`SELECT * FROM BOOKS`,
	)
	if err != nil {
		panic(err)
	}

	// 処理が終わったらカーソルを閉じる
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string

		// カーソルから値を取得
		if err := rows.Scan(&id, &title); err != nil {
			log.Fatal("rows.Scan()", err)
			return
		}

		fmt.Printf("id: %d, title: %s\n", id, title)
	}
}
