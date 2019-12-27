package postgres


import pb "github.com/uzdigital12345/golang/models_grpc_task_list/proto"

type TaskListInterface interface {
	Add(t *pb.Task) error
	UpdateTitle(id int64, title string) error
	Delete(id int64) error
	MakeDone(id int64) error
	GetTask(id int64) (*pb.Task,error)
	GetUnfinishedTasks() ([]*pb.Task,error)
	GetOverdueTasks() ([]*pb.Task,error)
	GetAll() ([]*pb.Task,error)
}
