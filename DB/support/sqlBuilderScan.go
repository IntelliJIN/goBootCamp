package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/doug-martin/goqu.v4"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool // Valid is true if Time is not NULL
}

// Scan implements the Scanner interface.
func (nt *NullTime) Scan(value interface{}) error {
	nt.Time, nt.Valid = value.(time.Time)
	return nil
}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func main() {
	sqDb, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
	db := goqu.New("sqlite3", sqDb)

	defer db.Db.Close()

	// START1 OMIT
	type Users struct {
		UId         int
		Username    sql.NullString
		Depart_Name sql.NullString
		Purpose     sql.NullString
		Ability     sql.NullString
		Created     NullTime
	}

	var users []Users
	err = db.From("users").ScanStructs(&users)
	checkErr(err)
	for _, user := range users {
		fmt.Printf("%v\n", user)
	}
	var use []Users
	err = db.From("users").Select("username").ScanStructs(&use)
	checkErr(err)

	for _, user := range use {
		fmt.Printf("\n%+v\n", user)
	}
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
