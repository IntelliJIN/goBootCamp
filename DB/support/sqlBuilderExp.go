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

	sqlString, _, _ := db.From("test").Where(
		goqu.Or(
			goqu.I("a").Gt(10),
			goqu.And(
				goqu.I("b").Eq(100),
				goqu.I("c").Neq("test"),
			),
		),
	).ToSql()
	fmt.Println(sqlString)
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
