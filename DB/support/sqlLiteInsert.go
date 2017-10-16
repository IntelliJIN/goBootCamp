package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS
		userinfo (uid INTEGER PRIMARY KEY AUTOINCREMENT,
		username VARCHAR(64) NULL,
		departname VARCHAR(64) NULL,
		created DATE NULL
	);`
	_, err = db.Exec(createTable)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) VALUES(?,?,?)")
	checkErr(err)

	for i := 0; i < 10; i++ {
		_, err := stmt.Exec("Mr Smith", "Matrix Police", "2001-12-09")
		checkErr(err)
	}
	count := getCount(db)
	fmt.Printf("System Error. Agent Smith had copied himself %v times\n", count-1)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getCount(db *sql.DB) (count int) {
	err := db.QueryRow("SELECT count(uid) FROM userinfo").Scan(&count)
	checkErr(err)
	return
}
