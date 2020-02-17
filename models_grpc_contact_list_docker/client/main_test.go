package client

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pos "github.com/uzdigital12345/golang/models_grpc/postgres"
	pb "github.com/uzdigital12345/golang/models_grpc/proto"
)

const (
	username = "postgres"
	password = "123"
	dbname   = "testdb"
	dbname2  = "task_list"
	bufSize  = 1024 * 1024
)

type service struct {
	ps pos.ContactManagerInterface
}

var lis *bufconn.Listener

func TestMain(t *testing.T) {

	psqlInfo := fmt.Sprintf(" user=%s password =%s dbname = %s sslmode=disable", username, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		t.Error("sqlx connection to postgres error")
		return
	}

	if err != nil {
		log.Fatal("Error listening...", err)
	}
	srv := grpc.NewServer()
	nw := pos.NewContactManagerInterface(db)

	lis := bufconn.Listen(bufSize)

	pb.RegisterContactManagerInterfaceServer(srv, &service{nw})

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatal("Error while listening: %v ",err)
		}
	} ()
}


