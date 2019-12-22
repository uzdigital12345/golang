package models

import "github.com/jinzhu/gorm"

type Contact struct {
	gorm.Model
	Age int
	Name string
	Gender string
	PhoneNumber string
}


type ContactManagerInterface interface {
	Add(c Contact) error
	UpdateName(id uint,name string) error
	Delete(i uint) error
	ListAll() error
	GetAll() ([]Contact,error)
	GetPaging(page int,limit int) ([]Contact,error)
}