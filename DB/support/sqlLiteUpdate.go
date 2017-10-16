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

func selectUser(db *sql.DB, uid int) (username string) {
	err := db.QueryRow("SELECT username FROM userinfo WHERE uid=?", uid).Scan(&username)
	checkErr(err)
	return
}

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	defer db.Close()

	createTable := `CREATE TABLE IF NOT EXISTS userinfo (uid INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(64) NULL,
	departname VARCHAR(64) NULL,
	created DATE NULL
	);`
	_, err = db.Exec(createTable)
	checkErr(err)

	userID := getUserID(db)

	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err := stmt.Exec("Original Mr Smith", userID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("Rows Affected", affect)
	fmt.Println("Original Mr Smith Found and Upgraded")

	fmt.Println("Original Mr Smith Found and Upgraded. New name: ", selectUser(db, userID))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
