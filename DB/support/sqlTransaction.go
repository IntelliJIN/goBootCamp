package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func getUserID(db *sql.DB) (uid int) {
	err := db.QueryRow("SELECT uid FROM userinfo").Scan(&uid)
	checkErr(err)
	return
}

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	fmt.Println("Opened Connections: ", db.Stats().OpenConnections)

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

	uid := getUserID(db)

	trashSQL, err := db.Prepare("update userinfo set username='Y',created=datetime() where uid=?")
	checkErr(err)

	tx, err := db.Begin()
	checkErr(err)

	_, err = tx.Stmt(trashSQL).Exec(uid)
	if err != nil {
		fmt.Println("Doing Rollback")
		tx.Rollback()
	} else {
		tx.Commit()
		fmt.Println("Transaction Committed")
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
