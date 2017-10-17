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

	fmt.Println("Opened Connections: ", db.Stats().OpenConnections)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	checkErr(err)

	defer db.Close()
	fmt.Println("Opened Connections: ", db.Stats().OpenConnections)
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
