package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Email struct {
	gorm.Model
	Email  string
	UserID uint `gorm:"ForeignKey:UId"`
}

func (e *Email) TableName() string {
	return "emails"
}

func (e *Email) BeforeSave() (err error) {
	fmt.Println("trigger on before save")
	return
}

func main() {
	// START1 OMIT
	db, err := gorm.Open("sqlite3", "./foo.db")
	checkErr(err)
	defer db.Close()

	db.AutoMigrate(&Email{})

	db.Create(&Email{Email: "neo@matrix.com", UserID: 2})

	FindByID(db, 1)
	// STOP1 OMIT
}

func checkErr(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
		panic(err)
	}
}

func FindByID(db *gorm.DB, id uint) {
	email := Email{}
	err := db.Find(&email, id).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Record not found", id)
	} else {
		checkErr(err)
	}
	fmt.Printf("PrintByID: %+v, data: %+v\n", id, email)
}
