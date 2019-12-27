package models

import (
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"

	pos "github.com/uzdigital12345/golang/models_grpc_task_list/postgres"
	pb "github.com/uzdigital12345/golang/models_grpc_task_list/proto"
)

var (
	a        pos.TaskListInterface
	client   pb.TaskListManagerInterfaceClient
	task     *pb.Task
	getTasks []*pb.GetTasks
	err      error
)

func TestSqlxDB_Add(t *testing.T) {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		t.Error("Connection error : ", err)
	}

	client = pb.NewTaskListManagerInterfaceClient(conn)

	task1 := &pb.Task{Assignee: "Abdurahmon", Title: "backend", Deadline: "2020.01.01", Done: false}
	task2 := &pb.Task{Assignee: "Doniyor", Title: "frontend", Deadline: "2020.01.15", Done: false}
	task3 := &pb.Task{Assignee: "Jahongir", Title: "testing system", Deadline: "2020.01.15", Done: false}
	task4 := &pb.Task{Assignee: "Temur", Title: "design", Deadline: "2019.12.28", Done: true}

	task, err = client.Add(context.Background(), task1)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Added item : %+v\n", task)

	task, err = client.Add(context.Background(), task2)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Added item : %+v\n", task)

	task, err = client.Add(context.Background(), task3)
	if err != nil {
		t.Error("Can't Add", err)
	}
	fmt.Printf("Added item : %+v\n", task)

	task, err = client.Add(context.Background(), task4)
	if err != nil {
		t.Error("Can't Add in testing ... ", err)
	}
	fmt.Printf("Added item : %+v\n", task)
	if err != nil {
		t.Error("Can't Add", err)
	}

}

func TestSqlxDB_UpdateTitle(t *testing.T) {
	task12 := &pb.UpdateTaskRequest{Id: 1, Title: "frontend"}
	task, err = client.UpdateTitle(context.Background(), task12)
	if err != nil {
		t.Error("Can't UpdateTitle", err)
	}
	fmt.Printf("Updated item : %+v\n", task)
}

func TestSqlxDB_Delete(t *testing.T) {
	task, err = client.Delete(context.Background(), &pb.IdRequest{Id: 1})
	if err != nil {
		t.Error("Can't delete", err)
	}
}

func TestSqlxDB_GetTask(t *testing.T) {
	task, err = client.GetTask(context.Background(), &pb.IdRequest{Id: 1})
	if err != nil {
		t.Error("Can't GetTask", err)
	}
	fmt.Printf("Get Task : %+v\n", task)
}

func TestSqlxDB_GetUnfinishedTasks(t *testing.T) {
	a, err := client.GetUnfinishedTasks(context.Background(), &pb.EmptyMessage{})
	if err != nil {
		t.Error("Can't GetUnfinishedTask", err)
	}
	fmt.Printf("Get Unfinished Tasks : %+v\n", a)
}

func TestSqlxDB_GetOverdueTasks(t *testing.T) {
	a, err := client.GetOverdueTasks(context.Background(), &pb.EmptyMessage{})
	if err != nil {
		t.Error("Can't GetOverdueTasks", err)
	}
	fmt.Printf("Get Overdue Tasks : %+v\n", a)
}

func TestSqlxDB_GetAll(t *testing.T) {

	a, err := client.GetAll(context.Background(), &pb.EmptyMessage{})
	if err != nil {
		t.Error("Can't GetTask", err)
	}
	fmt.Printf("Get Task : %+v\n", a)
}
