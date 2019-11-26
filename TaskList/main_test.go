package task

import (
	"fmt"
	"testing"
)
var (
	tm *TaskListManager
    t1 = TaskList{0, "Hojiakbar", "design", "2019-12-01", false}
    t2 = TaskList{1, "Temur", "backend", "2019-12-05", true}
	t3 = TaskList{2, "Sardor", "frontend", "2019-10-31", false}
	t4 = TaskList{3, "Akbar", "backend", "2019-09-01", true}
)

func TestTaskListManager_Add(t *testing.T) {
	tm = NewTaskListManager()
	tm.Add(t1)
	tm.Add(t2)
	tm.Add(t3)
	tm.Add(t4)
	fmt.Println("List All")
	tm.ListAll()
	if tm.tasks[0].assignee!="Hojiakbar" {
		t.Error("t1 is not added")
	}
	if tm.tasks[1].assignee!="Temur" {
		t.Error("t2 is not added")
	}
}

func TestTaskListManager_UpdateTitle(t *testing.T) {
	tm.UpdateTitle(t2,"updated")
	if tm.tasks[t2.id].title!="updated" {
		t.Error("It is not updated")
	}
}

func TestTaskListManager_Delete(t *testing.T) {
	tm.Delete(t2)
	if tm.tasks[1].assignee == "Temur" {
		t.Error("It is not deleted")
	}
}
func TestTaskListManager_MakeDone(t *testing.T) {
	tm.MakeDone(t1)
	if !tm.tasks[0].done {
		t.Error("It is not changed")
	}
}