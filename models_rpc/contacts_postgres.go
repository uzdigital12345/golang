package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

const (
	username = "postgres"
	password = "123"
	dbname = "testdb"
	dbname2 = "task_list"
	port = 4040
)

var err error

type Sqlx struct {
	fb *sqlx.DB
}

func NewContactManagerInterface() (ContactManagerInterface, error) {

	cm := Sqlx{}

	psqlInfo := fmt.Sprintf(" user=%s password =%s dbname = %s sslmode=disable",username,password,dbname)

	cm.fb, err = sqlx.Connect("postgres",psqlInfo)


	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	cm.fb, err = sqlx.Connect("postgres",psqlInfo)

	return &cm, nil
}
func (s *Sqlx) Add(c Contact,reply *Contact) error {

	insertionQuery:="insert into users (age, name, gender, number) values ($1, $2, $3, $4);"
	_, err := s.fb.Exec(insertionQuery, c.Age, c.Name, c.Gender, c.PhoneNumber)

	if err != nil {
		return err
	}

	*reply = c

	return nil

}

func (s *Sqlx) Update(c Contact,reply *Contact) error {

	updatingQuery:="update users set age=$1,name=$2,gender=$3,number=$4 where id=$5;"

	_, err = s.fb.Exec(updatingQuery,c.Age,c.Gender,c.Name, c.PhoneNumber,c.Id)


	if err!=nil {
		return err
	}

	*reply = c

	return nil
}

func (s *Sqlx) Delete(contacts2 Contact,reply *Contact) error {

	deletingQuery := "delete  from users where id=$1;"

	_,err = s.fb.Exec(deletingQuery,contacts2.Id)

	if err != nil {
		return err
	}

	*reply = contacts2

	return nil
}

func (s *Sqlx) ListAll(empty string,reply *[]Contact) error{
	var contacts []Contact
	listAllQuery:="select * from users;"

	err := s.fb.Select(&contacts,listAllQuery)

	if err != nil {
		return err
	}

	for index,_:=range contacts {
		fmt.Println(contacts[index])
	}
	*reply = contacts
	return nil

}

func (s *Sqlx) GetAll(empty string,reply *[]Contact) error {
	var contacts []Contact
	listAllQuery:="select * from users;"

	err := s.fb.Select(&contacts,listAllQuery)

	if err != nil {
		return err
	}
	*reply = contacts
	return nil

}

func main() {
	cm, err := NewContactManagerInterface()
	if err != nil {
		log.Fatalln("Error registering ...",err)
	}

	err = rpc.Register(cm)

	if err != nil {
		log.Fatalln("Error registering ...",err)
	}

	rpc.HandleHTTP()

	a :=fmt.Sprintf(":%d",port)
	listener, err := net.Listen("tcp",a)

	if err != nil {
		log.Fatal("Error listener",err)
	}

	log.Printf("serving rpc on port %d",port)

	err = http.Serve(listener,nil)

	if err != nil {
		log.Fatal("Error serving : ",err)
	}

}
