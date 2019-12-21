package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	Db ,err := gorm.Open("postgres" , "user=postgres password=123 dbname = gorm sslmode=disable")
	if err != nil {
		panic(err.Error())
	}

	defer Db.Close()
	Db.DropTableIfExists(&Owner{},&Book{},&Author{})
	Db.CreateTable(&Owner{},&Book{},&Author{})

}

type Owner struct {
	gorm.Model
	FirstName string
	LastName string
	Books []Book

}

type Book struct {
	gorm.Model
	Name string
	Published time.Time
	OwnerId uint `sql:"index"`
	authors []Author `gorm:"many2many:books_authors"`
}

type Author struct {
	gorm.Model
	FirstName string
	LastName string
}

