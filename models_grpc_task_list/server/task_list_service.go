package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pos "github.com/uzdigital12345/golang/models_grpc_task_list/postgres"
	pb "github.com/uzdigital12345/golang/models_grpc_task_list/proto"
)

const (
	username = "postgres"
	password = "123"
	dbname   = "task_list"
)

var (
	ts  pos.TaskListInterface
	err error
)

type service struct {
	ts pos.TaskListInterface
}

func (s *service) Add(ctx context.Context, t *pb.Task) (*pb.Task, error) {
	err = s.ts.Add(t)
	if err != nil {
		log.Fatal("Can't Add in task_list_service", err)
		return nil, err
	}
	return t, err
}

func (s *service) UpdateTitle(ctx context.Context, t *pb.UpdateTaskRequest) (*pb.Task, error) {
	err = s.ts.UpdateTitle(t.GetId(), t.GetTitle())
	if err != nil {
		log.Fatal("Can't Update in task_list_service", err)
		return nil, err
	}
	t2, err := ts.GetTask(t.GetId())
	if err != nil {
		log.Fatal("Can't get Task in Update in task_list_service", err)
		return nil, err
	}
	return t2, nil
}

func (s *service) Delete(ctx context.Context, t *pb.IdRequest) (*pb.Task, error) {
	err = s.ts.Delete(t.GetId())
	if err != nil {
		log.Fatal("Can't Delete in task_list_service", err)
		return nil, err
	}
	t2, err := ts.GetTask(t.GetId())
	if err != nil {
		log.Fatal("Can't get Task in Delete in task_list_service", err)
		return nil, err
	}
	return t2, nil
}

func (s *service) MakeDone(ctx context.Context, t *pb.IdRequest) (*pb.Task, error) {
	err = s.ts.MakeDone(t.GetId())
	if err != nil {
		log.Fatal("Can't MakeDone in task_list_service", err)
		return nil, err
	}
	t2, err := ts.GetTask(t.GetId())
	if err != nil {
		log.Fatal("Can't get Task in MakeDone in task_list_service", err)
		return nil, err
	}
	return t2, nil
}

func (s *service) GetTask(ctx context.Context, t *pb.IdRequest) (*pb.Task, error) {

	t2, err := ts.GetTask(t.GetId())
	if err != nil {
		log.Fatal("Can't get Task  in task_list_service", err)
		return nil, err
	}
	return t2, nil
}

func (s *service) GetOverdueTasks(ctx context.Context, empty *pb.EmptyMessage) (*pb.GetTasks, error) {


	t2, err := ts.GetOverdueTasks()
	if err != nil {
		log.Fatal("Can't get Overdue Task  in task_list_service", err)
		return nil, err
	}
	return &pb.GetTasks{
		Tasks: t2,
	}, nil
}

func (s *service) GetUnfinishedTasks(ctx context.Context, empty *pb.EmptyMessage) (*pb.GetTasks, error) {

	t2, err := ts.GetUnfinishedTasks()
	if err != nil {
		log.Fatal("Can't get Unfinished Task  in task_list_service", err)
		return nil, err
	}
	return &pb.GetTasks{
		Tasks: t2,
	}, nil
}

func (s *service) GetAll(ctx context.Context, empty *pb.EmptyMessage) (*pb.GetTasks, error) {

	t2, err := ts.GetAll()
	if err != nil {
		log.Fatal("Can't get Get All Task  in task_list_service", err)
		return nil, err
	}
	return &pb.GetTasks{
		Tasks: t2,
	}, nil
}

func main() {
	psqlInfo := fmt.Sprintf(" user=%s password =%s dbname = %s sslmode=disable", username, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connection database ...", err)
	}
	ts = pos.NewTaskList(db)

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Error listening...", err)
	}
	srv := grpc.NewServer()
	pb.RegisterTaskListManagerInterfaceServer(srv, &service{ts})
	reflection.Register(srv)
	fmt.Println("Server is running... Port : 4040")
	err = srv.Serve(listener)
	if err != nil {
		log.Fatal("Error serving. .. ")
	}
}
