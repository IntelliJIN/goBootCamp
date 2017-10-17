package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// START1 OMIT
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	defer db.Close()
	userID := 1

	// delete
	stmt, err := db.Prepare("delete from users where uid=?")
	checkErr(err)

	res, err := stmt.Exec(userID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("Rows Affected", affect)
	fmt.Println("System Info. Original Mr Smith Found and deleted. System safe now.")
	// STOP1 OMIT
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}