package models

type Contact struct {
	Id uint
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
}