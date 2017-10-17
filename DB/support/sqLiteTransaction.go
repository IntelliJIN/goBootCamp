package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func getUserID(db *sql.DB) (uid int) {
	err := db.QueryRow("SELECT uid FROM users").Scan(&uid)
	checkErr(err)
	return
}

func main() {
	//START1 OMIT
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	defer db.Close()

	uid := getUserID(db)

	trashSQL, err := db.Prepare("update users set username=?, created=? where uid=?")
	checkErr(err)

	tx, err := db.Begin()
	checkErr(err)

	_, err = tx.Stmt(trashSQL).Exec("Neo", time.Now(), uid)
	if err != nil {
		fmt.Println("Doing Rollback")
		tx.Rollback()
	} else {
		tx.Commit()
		fmt.Println("Transaction Committed")
	}
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
