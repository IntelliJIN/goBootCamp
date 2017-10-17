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

	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	var uid int
	var username, department, purpose string
	var ability sql.NullString
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &purpose, &ability, &created)
		checkErr(err)
		fmt.Println("UserID: ", uid, "UserName: ", username, "Department: ", department,
			"Ability: ", ability, "Purpose: ", purpose, "Created: ", created)
	}

	rows.Close() //good habit to close
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
