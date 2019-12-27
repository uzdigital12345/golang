package postgres

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	pb "github.com/uzdigital12345/golang/models_grpc_task_list/proto"
)

type sqlxDB struct {
	connectDB *sqlx.DB
}

var (
	err      error
	deadline time.Time
)

func NewTaskList(db *sqlx.DB) TaskListInterface {
	return &sqlxDB{connectDB: db}
}

func (s *sqlxDB) Add(t *pb.Task) error {
	insertionQuery := "insert into task_list_info (assignee, title, deadline, done) values ($1, $2, $3, $4);"

	_, err := s.connectDB.Exec(insertionQuery, t.Assignee, t.Title, t.Deadline, t.Done)

	if err != nil {
		log.Fatal("Error adding in task_list_postgres", err)
		return err
	}

	return nil
}

func (s *sqlxDB) UpdateTitle(id int64, title string) error {
	updatingQuery := "update task_list_info set title=$1 where id=$2;"

	_, err := s.connectDB.Exec(updatingQuery, title, id)

	if err != nil {
		log.Fatal("Can't update in postgres ..", err)
		return err
	}

	return nil
}

func (s *sqlxDB) Delete(id int64) error {
	deletingQuery := "delete from task_list_info where id=$1;"

	_, err := s.connectDB.Exec(deletingQuery, id)

	if err != nil {
		log.Fatal("Can't delete in postgres", err)
		return err
	}
	return nil
}

func (s *sqlxDB) MakeDone(id int64) error {
	makeDoneQuery := "update task_list_info set done=$1 where id=$2;"

	_, err := s.connectDB.Exec(makeDoneQuery, "true", id)

	if err != nil {
		log.Fatal("Can't make done in postgres", err)
		return err
	}
	return nil
}

func (s *sqlxDB) GetTask(id int64) (*pb.Task, error) {
	listTaskQuery := "select * from task_list_info where id=$1;"

	rows, err := s.connectDB.Queryx(listTaskQuery, id)

	if err != nil {
		log.Fatal("Can't get task in postgres", err)
		return nil, err
	}
	var ts pb.Task
	for rows.Next() {
		err = rows.StructScan(&ts)
		if err != nil {
			log.Fatal("Can't scan struct in GetTask", err)
			return nil, err
		}
	}

	return &ts, nil
}

func (s *sqlxDB) GetUnfinishedTasks() ([]*pb.Task, error) {
	listUnfinishedTasksQuery := "select * from task_list_info where deadline < now() and done = false ;"

	rows, err := s.connectDB.Queryx(listUnfinishedTasksQuery)
	if err != nil {
		log.Fatal("Can't get UnfinishedTasks in postgres", err)
		return nil, err
	}
	var t []*pb.Task
	for rows.Next() {
		var ts pb.Task
		err = rows.StructScan(&ts)
		if err != nil {
			log.Fatal("Can't scan to struct in GetUnfinishedTasks", err)
			return nil, err
		}
		t = append(t, &ts)
	}
	return t, nil
}

func (s *sqlxDB) GetOverdueTasks() ([]*pb.Task, error) {
	listOverdueTasksQuery := "select * from task_list_info where deadline > now() and done=false;"

	rows, err := s.connectDB.Queryx(listOverdueTasksQuery)

	if err != nil {
		log.Fatal("Error in GetOverdueTasks in postgres", err)
		return nil, err
	}
	var t []*pb.Task
	for rows.Next() {
		var ts pb.Task
		err := rows.StructScan(&ts)
		if err != nil {
			log.Fatal("Can't scan to struct in GetOverdueTasks", err)
			return nil, err
		}
		t = append(t, &ts)
	}
	return t, nil
}

func (s *sqlxDB) GetAll() ([]*pb.Task, error) {
	getAllQuery := "select * from task_list_info;"

	rows, err := s.connectDB.Queryx(getAllQuery)

	if err != nil {
		log.Fatal("Error in GetAll in postgres", err)
		return nil, err
	}

	var t []*pb.Task

	for rows.Next() {
		var ts pb.Task
		err = rows.StructScan(&ts)
		if err != nil {
			log.Fatal("Can't scan to struct in GetAll", err)
			return nil, err
		}
		t = append(t, &ts)
	}
	return t, nil
}
