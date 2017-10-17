package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	// START1 OMIT
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO users(username, depart_name, purpose ,created) values(?,?,?,?)")
	checkErr(err)

	for i := 0; i < 10; i++ {
		_, err := stmt.Exec(
			"Agent Smith Copy",
			"Matrix Police",
			"Protect the Matrix",
			time.Now())
		checkErr(err)
	}
	count := getCount(db)
	fmt.Printf("System Error. Agent Smith have olready %v copies\n", count-1)
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// START2 OMIT
func getCount(db *sql.DB) (count int) {
	err := db.QueryRow("SELECT count(uid) FROM users").Scan(&count)
	checkErr(err)
	return
}

// STOP2 OMIT
