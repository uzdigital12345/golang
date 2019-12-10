package models

type Contact struct {
	Id int `db:"id"`
	Age int `db:"age"`
	Name string `db:"name"`
	Gender string `db:"gender"`
	PhoneNumber int `db:"number"`
}

type ContactManagerInterface interface {
	Add(c Contact) error
	Update(i int,c Contact) error
	Delete(i int) error
	ListAll() error
	GetAll() ([]Contact,error)
}