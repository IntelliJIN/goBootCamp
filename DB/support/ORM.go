package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type User struct {
	UId        int       `gorm:"primary_key;column:uid"`
	Username   string    `gorm:"column:username"`
	DepartName string    `gorm:"column:depart_name"`
	Purpose    string    `gorm:"column:purpose"`
	Ability    string    `gorm:"column:ability"`
	Created    time.Time `gorm:"column:created"`
}

func main() {
	// START1 OMIT
	db, err := gorm.Open("sqlite3", "./foo.db")
	checkErr(err)
	defer db.Close()

	db.Create(&User{Username: "R2D2", DepartName: "Robots", Created: time.Now()})

	var user User
	db.First(&user, 11)
	fmt.Println("User Smith: ", user)

	user = *new(User)
	db.First(&user, "username = ?", "R2D2")
	fmt.Println("User R2D2: ", user)

	db.Model(&user).Update("ability", "Fly").Where(&User{Username: "Trinity"})
	fmt.Println("Updated User Trinity: ", user)

	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
