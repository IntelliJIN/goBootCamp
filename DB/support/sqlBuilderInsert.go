package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/doug-martin/goqu.v4"
	"time"
)

func main() {
	sqDb, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	db := goqu.New("sqlite3", sqDb)

	defer db.Db.Close()

	// START1 OMIT

	users := []goqu.Record{
		{"username": "Morpheus", "depart_name": "Rebels", "created": time.Now()},
		{"username": "Trinity", "depart_name": "Rebels", "created": time.Now()},
		{"username": "Tank", "depart_name": "Rebels", "created": time.Now()},
	}
	_, err = db.From("users").Insert(users).Exec()
	checkErr(err)
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
