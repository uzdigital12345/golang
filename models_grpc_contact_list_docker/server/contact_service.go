package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pos "github.com/uzdigital12345/golang/models_grpc/postgres"
	pt "github.com/uzdigital12345/golang/models_grpc/proto"
)

const (
	username = "postgres"
	password = "123"
	dbname   = "testdb"
	dbname2  = "task_list"
)

var (
	nw  pos.ContactManagerInterface
	err error
)

type service struct {
	ps pos.ContactManagerInterface
}

func (s *service) Add(ctx context.Context, ct *pt.Contact) (*pt.Contact, error) {

	err = s.ps.Add(ct)
	if err != nil {
		return nil, err
	}
	addedContact,err := s.ps.GetContact(ct.GetId())
	if err != nil {
		return nil, err
	}
	return addedContact, nil
}

func (s *service) Update(ctx context.Context, ct *pt.Contact) (*pt.Contact, error) {

	err = s.ps.Update(ct.GetId(), ct)
	if err != nil {
		return nil, err
	}
	updatedContact,err := s.ps.GetContact(ct.GetId())
	if err != nil {
		return nil, err
	}
	return updatedContact, nil

}

func (s *service) Delete(ctx context.Context, ct *pt.Contact) (*pt.Contact, error) {

	err = s.ps.Delete(ct.GetId())
	if err != nil {
		return nil, err
	}
	deletedContact,err := s.ps.GetContact(ct.GetId())
	if err != nil {
		return nil, err
	}
	return deletedContact, nil
}

func (s *service) GetAll(ctx context.Context, emt *empty.Empty) (*pt.GetContacts, error) {
	gt, err := s.ps.GetAll()
	if err != nil {
		return nil, err
	}
	return &pt.GetContacts{
		Contacts: gt,
	}, nil
}

func main() {

	psqlInfo := fmt.Sprintf(" user=%s password =%s dbname = %s sslmode=disable", username, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connection database ...", err)
	}

	nw = pos.NewContactManagerInterface(db)

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Error listening...", err)
	}
	srv := grpc.NewServer()
	pt.RegisterContactManagerInterfaceServer(srv, &service{nw})
	reflection.Register(srv)
	fmt.Println("Server is running... Port : 4040")
	err = srv.Serve(listener)
	if err != nil {
		log.Fatal("Error serving. .. ")
	}

}
