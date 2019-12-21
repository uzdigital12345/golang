package models

import (
	"testing"
)

var (
	a TaskListInterface
)

func TestNewTaskList(t *testing.T) {
	_, err = NewTaskList()
	if err != nil {
		t.Error("NewTaskList cannot be created")
	}
}

func TestGormDB_Add(t *testing.T) {
	a, err = NewTaskList()
	t1 := TaskList{6,"Abdurahmon","backend","2020.01.01",false}
	t2 := TaskList{7,"Doniyor","frontend","2020.01.15",false}
	t3 := TaskList{8,"Jahongir","testing system","2020.01.15",false}
	t4 := TaskList{9,"Temur","design","2019.12.28",true}
	if err != nil {
		t.Error("Can't Add", err)
	}
	err = a.Add(t1)
	err = a.Add(t2)
	err = a.Add(t3)
	err = a.Add(t4)

	if err != nil {
		t.Error("Can't Add", err)
	}
}

func TestGormDB_UpdateTitle(t *testing.T) {
	a.UpdateTitle(7,"designer")
	if err != nil {
		t.Error("Can't update title")
	}
}

func TestGormDB_Delete(t *testing.T) {
	if err != nil {
		t.Error("Can't delete")
	}
	err = a.Delete(3)
}

func TestGormDB_ListTask(t *testing.T) {
	err = a.ListTask(2)
	if err != nil {
		t.Error("Error while listing task: ", err)
	}
}

func TestGormDB_ListAll(t *testing.T) {
	err = a.ListAll()
	if err != nil {
		t.Error("Can't execute list all")
	}
}


