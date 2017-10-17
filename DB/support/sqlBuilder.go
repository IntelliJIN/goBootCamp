package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/doug-martin/goqu.v4"
)

func main() {
	// START1 OMIT
	sqDb, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	db := goqu.New("sqlite3", sqDb)

	defer db.Db.Close()

	sqlString, _, _ := db.From("users").Where(goqu.Ex{
		"id": 5,
	}).ToSql()
	fmt.Println(sqlString)

	//prepared sql
	sqlString, args, _ := db.From("users").
		Prepared(true).
		Where(goqu.Ex{
			"id": 10,
		}).ToSql()
	fmt.Println(sqlString, args)
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
