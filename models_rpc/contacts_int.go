package main

type Contact struct {
	Id int `db:"id"`
	Age int `db:"age"`
	Name string `db:"name"`
	Gender string `db:"gender"`
	PhoneNumber int `db:"number"`
}

type ContactManagerInterface interface {
	Add(c Contact,reply *Contact) error
	Update(c Contact,reply *Contact) error
	Delete(contacts2 Contact,reply *Contact) error
	ListAll(empty string,reply *[]Contact) error
	GetAll(empty string,reply *[]Contact) error
}