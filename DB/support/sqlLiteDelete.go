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

	createTable := `CREATE TABLE IF NOT EXISTS userinfo (uid INTEGER PRIMARY KEY AUTOINCREMENT,
	username VARCHAR(64) NULL,
	departname VARCHAR(64) NULL,
	created DATE NULL
	);`
	_, err = db.Exec(createTable)
	checkErr(err)

	userID := getUserID(db)

	// delete
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err := stmt.Exec(userID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("Rows Affected", affect)
	fmt.Println("System Info. Original Mr Smith Found and deleted. System safe now.")
}
