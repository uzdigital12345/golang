package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type sqlxDB struct{
	connectDB *sqlx.DB
}

func  NewTaskList() (TaskListInterface,error){
	cm := sqlxDB{}
	psqlInfo := fmt.Sprintf(" user=%s dbname=%s password=%s",username,dbname2,password)
	cm.connectDB,err = sqlx.Connect("postgres",psqlInfo)
	if err!=nil {
		fmt.Println(err)
		return nil,err
	}
	return &cm,nil
}

func (s sqlxDB) Add(t TaskList) error {
	insertionQuery:="insert into task_list_info (assignee, title, deadline, done) values ($1, $2, $3, $4);"

	_, err := s.connectDB.Exec(insertionQuery, t.Assignee, t.Title, t.Deadline, t.Done)

	if err != nil {
		return err
	}

	return nil
}

func (s sqlxDB) UpdateTitle(id int,title string) error {
	updatingQuery := "update task_list_info set title=$1 where id=$2;"

	_, err := s.connectDB.Exec(updatingQuery,title,id)

	if err != nil {
		fmt.Println("Can't update")
		return err
	}

	return nil
}

func (s sqlxDB) Delete(id int) error{
	deletingQuery := "delete from task_list_info where id=$1;"

	_,err := s.connectDB.Exec(deletingQuery,id)

	if err != nil {
		fmt.Println("Can't delete")
		return err
	}
	return nil
}

func(s sqlxDB) MakeDone(id int) error {
	makeDoneQuery := "update task_list_info set done=$1 where id=$2;"

	_,err :=s.connectDB.Exec(makeDoneQuery,"true",id)

	if err != nil {
		fmt.Println("Can't make done")
		return err
	}
	return nil
}

func (s sqlxDB) ListTask(id int) error {
	listTaskQuery := "select * from task_list_info where id=$1;"

	rows, err := s.connectDB.Queryx(listTaskQuery,id)

	if err != nil {
		fmt.Println("Can't print task list")
		return err
	}
	var ts TaskList
	for rows.Next() {
		err = rows.StructScan(&ts)
		if err != nil {
			fmt.Println("Can't scan struct")
			return err
		}
	}
	fmt.Println(ts)
	return nil
}

func (s sqlxDB) ListUnfinishedTasks() error {
	listUnfinishedTasksQuery := "select * from task_list_info where deadline < now() and done = false ;"

	rows,err := s.connectDB.Queryx(listUnfinishedTasksQuery)
	if err != nil {
		fmt.Println("Can't execute ListUnfinishedTasks")
		return err
	}
	var t []TaskList
	for rows.Next() {
		var ts TaskList
		err = rows.StructScan(&ts)
		if err != nil {
			fmt.Println("Can't scan to struct")
			return err
		}
		t = append(t,ts)
	}
	PrintSlice(t)
	return nil
}

func (s sqlxDB) ListOverdueTasks() error {
	listOverdueTasksQuery := "select * from task_list_info where deadline > now() done=false;"

	rows,err :=s.connectDB.Queryx(listOverdueTasksQuery)

	if err != nil {
		fmt.Println("Can't print overdue tasks")
		return err
	}
    var t []TaskList
	for rows.Next() {
		var ts TaskList
		err:=rows.StructScan(&ts)
		if err != nil {
			fmt.Println("Can't scan to struct")
			return err
		}
		t = append(t,ts)
	}
	PrintSlice(t)
	return nil
}

func (s sqlxDB) ListAll() error {
	listAllQuery := "select * from task_list_info;"

	rows,err :=s.connectDB.Queryx(listAllQuery)

	if err != nil {
		fmt.Println("Can't get list all")
		return err
	}

	var t []TaskList

	for rows.Next() {
		var ts TaskList
		err = rows.StructScan(&ts)
		if err != nil {
			fmt.Println("Can't scan to struct")
			return err
		}
		t = append(t, ts)
	}
	PrintSlice(t)
	return nil
}

func (s sqlxDB) GetAll() ([]TaskList, error) {
	getAllQuery := "select * from task_list_info;"

	rows,err :=s.connectDB.Queryx(getAllQuery)

	if err != nil {
		fmt.Println("Can't get all rows")
		return nil,err
	}

	var t []TaskList

	for rows.Next() {
		var ts TaskList
		err = rows.StructScan(&ts)
		if err != nil {
			fmt.Println("Can't scan to struct")
			return nil,err
		}
		t = append(t, ts)
	}
	PrintSlice(t)
	return t, nil
}

func PrintSlice(t []TaskList) {
	for i,_:=range t {
		fmt.Println(t[i])
	}
}




