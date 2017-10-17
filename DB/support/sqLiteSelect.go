package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
	// START1 OMIT
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	defer db.Close()

	var uid int
	var userName, departName, purpose, ability sql.NullString
	var created NullTime
	err = db.QueryRow("SELECT * FROM users").
		Scan(
			&uid,
			&userName,
			&departName,
			&purpose,
			&ability,
			&created)
	checkErr(err)
	fmt.Printf("UseID: %v\nUserName: %v\nDepartmentName: %v\n", uid, userName, departName)
	fmt.Printf("Purpose: %v\nAbility: %v\nJoined: %v", purpose, ability, created)
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
