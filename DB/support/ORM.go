package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

// START2 OMIT
type User struct {
	UId        int       `gorm:"primary_key;column:uid"`
	Username   string    `gorm:"column:username"`
	DepartName string    `gorm:"column:depart_name"`
	Purpose    string    `gorm:"column:purpose"`
	Ability    string    `gorm:"column:ability"`
	Created    time.Time `gorm:"column:created"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeSave() (err error) {
	fmt.Println("Before save")
	return
}

// STOP2 OMIT

func main() {
	// START1 OMIT
	db, err := gorm.Open("sqlite3", "./foo.db")
	checkErr(err)
	defer db.Close()

	if ok := db.HasTable("users"); !ok {
		fmt.Println("No such table")
	}

	FindByID(db, 1)
	FindByID(db, 1000)

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
		fmt.Println("ERROR :", err)
		panic(err)
	}
}

func FindByID(db *gorm.DB, id uint) {
	user := User{}
	err := db.Find(&user, id).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Record not found", id)
	} else {
		checkErr(err)
	}
	fmt.Printf("PrintByID: %+v, data: %+v\n", id, user)
}
