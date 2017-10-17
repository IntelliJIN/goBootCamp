package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func selectUser(db *sql.DB, uid int) (username, ability string) {
	err := db.QueryRow("SELECT username, ability FROM users WHERE uid=?", uid).Scan(&username, &ability)
	checkErr(err)
	return
}

func main() {
	// START1 OMIT
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	defer db.Close()

	stmt, err := db.Prepare("update users set username=?, ability=? where uid=?")
	checkErr(err)

	userID := 1
	res, err := stmt.Exec("Original Mr Smith", "Can replicate himself", userID)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	username, ability := selectUser(db, userID)
	fmt.Println("Rows Affected", affect)
	fmt.Println("Original Mr Smith Found and Upgraded")
	fmt.Println("New name: ", username)
	fmt.Println("New ability: ", ability)
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
