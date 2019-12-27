package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	pt "github.com/uzdigital12345/golang/models_grpc/proto"
)

const (
	username = "postgres"
	password = "123"
	dbname   = "testdb"
	dbname2  = "task_list"
)

var (
	err error
)

type Sqlx struct {
	db *sqlx.DB
}

func NewContactManagerInterface(db *sqlx.DB) ContactManagerInterface {
	return &Sqlx{db: db}
}

func (s *Sqlx) Add(c *pt.Contact) error {
	insertionQuery := "insert into users (age, name, gender, number) values ($1, $2, $3, $4);"

	_, err := s.db.Exec(insertionQuery, c.Age, c.Name, c.Gender, c.Number)

	if err != nil {
		return err
	}

	return nil

}

func (s *Sqlx) Update(i int64, c *pt.Contact) error {

	updatingQuery := "update users set age=$1,name=$2,gender=$3,number=$4 where id=$5;"

	_, err = s.db.Exec(updatingQuery, c.Age, c.Gender, c.Name, c.Number, i)

	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlx) Delete(i int64) error {

	deletingQuery := "delete from users where id=$1;"

	_, err = s.db.Exec(deletingQuery, i)

	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlx) GetAll() ([]*pt.Contact, error) {

	var c []*pt.Contact

	listAllQuery := "select * from users;"

	err := s.db.Select(&c, listAllQuery)

	if err != nil {
		return nil, err
	}

	return c, nil

}
