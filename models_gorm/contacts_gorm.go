package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	username = "postgres"
	password = "123"
	dbname = "contact_list"
	dbname2 = "task_list"
)


var err error

type GormDb struct {
	db *gorm.DB
}

func NewContactManagerInterface() (ContactManagerInterface,error) {

	nw := GormDb{}

	psqlInfo := fmt.Sprintf(" user=%s password =%s dbname = %s sslmode=disable",username,password,dbname)


	nw.db,err = gorm.Open("postgres",psqlInfo)
	nw.db.AutoMigrate(&Contact{})
	if err!=nil {
		panic(err.Error())
		return nil,err
	}
	nw.db.DropTableIfExists(&Contact{})
	nw.db.CreateTable(&Contact{})
	return &nw,nil
}

func (g *GormDb) TableName() string {
	return "contact_list_info"
}

func (g *GormDb) Add(c Contact) error {
	err = g.db.Create(&c).Error
	return Error(err)
}

func (g *GormDb) UpdateName(id uint,name string) error {
    err = g.db.Model(&Contact{}).Where("id = ?", id).Update("name", name).Error
	return err
}

func (g *GormDb) Delete(id uint) error {

	err = g.db.Where("id = ?",id).Delete(&Contact{}).Error
	return Error(err)
}

func (g *GormDb) ListAll() error{

	var contacts [] Contact
	err = g.db.Find(&contacts).Error
	return ErrorPrintContact(err,contacts)

}

func (g *GormDb) GetAll() ([]Contact,error) {

	var contacts [] Contact
	err = g.db.Find(&contacts).Error
	return contacts,Error(err)

}

func ErrorPrintContact(err error,contact []Contact) error {

	if err!=nil {
		return err
	}
	fmt.Println(contact)
	return nil

}
