package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	username = "postgres"
	password = "123"
	dbname = "testdb"
)


var err error

type Sqlx struct {
	fb *sqlx.DB
}

func NewContactManagerInterface() (ContactManagerInterface,error) {

	cm := Sqlx{}

	psqlInfo := fmt.Sprintf(" user=%s password =%s dbname = %s sslmode=disable",username,password,dbname)

	cm.fb, err = sqlx.Connect("postgres",psqlInfo)

	if err != nil {
		fmt.Println(err)
		return nil,err
	}

	return &cm, nil
}
func (s *Sqlx) Add(c Contact) error {
	insertionQuery:="insert into users (age, name, gender, number) values ($1, $2, $3, $4);"

	_, err := s.fb.Exec(insertionQuery, c.Age, c.Name, c.Gender, c.PhoneNumber)

	if err != nil {
		return err
	}

	return nil

}

func (s *Sqlx) Update(i int,c Contact) error {

	updatingQuery:="update users set age=$1,name=$2,gender=$3,number=$4 where id=$5;"

	_, err = s.fb.Exec(updatingQuery,c.Gender,c.Name, c.PhoneNumber,i)

	if err!=nil {
		return err
	}
	
	return nil
}

func (s *Sqlx) Delete(i int) error {

	deletingQuery:="delete * from users where id=$1;"

	_,err = s.fb.Exec(deletingQuery,i)

	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlx) ListAll() error{

	c := []Contact{}

	listAllQuery:="select * from users;"

	err := s.fb.Select(&c,listAllQuery)

	if err != nil {
		return err
	}

	for index,_:=range c {
		fmt.Println(c[index])
	}

	return nil

}

func (s *Sqlx) GetAll() ([]Contact,error) {

	c := []Contact{}

	listAllQuery:="select * from users;"

	err := s.fb.Select(&c,listAllQuery)

	if err != nil {
		return nil,err
	}

	return c,nil

}
