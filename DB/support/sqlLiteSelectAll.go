package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
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

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("User ID: ", uid, "User Name: ", username, "Department: ", department, "Created: ", created)
	}

	rows.Close() //good habit to close
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
