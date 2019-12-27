package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	// "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	pb "github.com/uzdigital12345/golang/models_grpc/proto"
)

const (
	port = 4040
)

type Sqlx struct {
	fb *sqlx.DB
}

var (
	err    error
	client pb.ContactManagerInterfaceClient
)

func TestSqlx_Add(t *testing.T) {

	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		t.Error("Connection error : ", err)
	}

	client = pb.NewContactManagerInterfaceClient(conn)
	contact1 := &pb.Contact{
		Age:         "19",
		Name:        "Enver",
		Gender:      "Male",
		Number: "95218125",
	}
	contact2 := &pb.Contact{
		Id:          2,
		Age:         "15",
		Name:        "Abdurahmon",
		Gender:      "Male",
		Number: "941563245",
	}
	contact3 := &pb.Contact{
		Id:          3,
		Age:         "85",
		Name:        "Boboy",
		Gender:      "Male",
		Number: "9415625315",
	}
	contact4 := &pb.Contact{
		Id:          3,
		Age:         "85",
		Name:        "Boboy",
		Gender:      "Male",
		Number: "9415625315",
	}

	contact, err := client.Add(context.Background(), contact1)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Added item : %+v\n", contact)

	contact, err = client.Add(context.Background(), contact2)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Added item : %+v\n", contact)

	contact, err = client.Add(context.Background(), contact3)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Added item : %+v\n", contact)

	contact, err = client.Add(context.Background(), contact4)
	if err != nil {
		t.Error("Can't Add in testing ... ", err)
	}
	fmt.Printf("Added item : %+v\n", contact)

}

func TestSqlx_Update(t *testing.T) {

	c := &pb.Contact{
		Age:         "15",
		Name:        "Akbar",
		Gender:      "Male",
		Number: "1581818",
	}

	contact, err := client.Update(context.Background(), c)
	if err != nil {
		t.Error("Can't update in testing...", err)
	}
	fmt.Printf("Updated item : %+v\n", contact)
}

func TestSqlx_Delete(t *testing.T) {
	//var b  Contact
	//c := Contact{ 5, 15, "Akbar", "male", 9898915}
	//
	//err = client.Call("Sqlx.Delete",c,&b)
	c := &pb.Contact{
		Age:         "15",
		Name:        "Akbar",
		Gender:      "Male",
		Number: "1581818",
	}

	b, err := client.Delete(context.Background(), c)
	if err != nil {
		t.Error("Can't delete in testing..", err)
	}
	fmt.Println(b)
}

func TestSqlx_GetAll(t *testing.T) {
	b, err := client.GetAll(context.Background(), &empty.Empty{})
	if err != nil {
		t.Error("Can't get all", err)
	}
	fmt.Println(b)
}
