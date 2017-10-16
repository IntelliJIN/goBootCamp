package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	fmt.Println("Opened Connections: ", db.Stats().OpenConnections)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	checkErr(err)

	defer db.Close()

	fmt.Println("Opened Connections: ", db.Stats().OpenConnections)

	createTable := `CREATE TABLE IF NOT EXISTS userinfo (uid INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(64) NULL,
	departname VARCHAR(64) NULL,
	created DATE NULL
	);`
	_, err = db.Exec(createTable)
	checkErr(err)

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("Mr Smith", "Matrix Police", "2001-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Mr Smith System ID:", id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
