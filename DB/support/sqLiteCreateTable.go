package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
// START1 OMIT
const createTable = `
	CREATE TABLE IF NOT EXISTS users (
	uid INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(64) NULL,
	depart_name VARCHAR(64) NULL,
	purpose VARCHAR(64) NULL,
	ability VARCHAR(64) NULL,
	created DATE NULL
);`
// STOP1 OMIT
func main() {
	// START2 OMIT
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	defer db.Close()

	_, err = db.Exec(createTable)
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO users(username, depart_name, purpose ,created) values(?,?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("Agent Smith", "Matrix Police", "Protect the Matrix", "1999-03-31")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Agent Smith System ID:", id)
	// STOP2 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
