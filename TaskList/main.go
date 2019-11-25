package main

import (
	"fmt"
	"time"
)

type TaskListManager struct {
	tasks []TaskList
}

func NewTaskListManager() *TaskListManager {
	return &TaskListManager{}
}

type TaskList struct {
	id       int
	assignee string
	title    string
	deadline string
	done     bool
}

func (tm *TaskListManager) Add(t TaskList) {
	tm.tasks = append(tm.tasks, t)
}

func (tm *TaskListManager) UpdateTitle(t TaskList, title string) {
	//tm.tasks[t.id].assignee = t.assignee
	tm.tasks[t.id].title = title
	//tm.tasks[t.id].deadline = t.deadline
	//tm.tasks[t.id].done = t.done
}

func (tm *TaskListManager) Delete(t TaskList) {
	tm.tasks = append(tm.tasks[:t.id], tm.tasks[t.id+1:]...)
}

func (tm *TaskListManager) MakeDone(t TaskList) {
	tm.tasks[t.id].done = true
}

func (tm *TaskListManager) ListTask(t TaskList) {
	fmt.Printf("%d.%s - %s\n",t.id+1,t.assignee,t.title)
}

func (tm *TaskListManager) ListUnfinishedTasks() {
	for i, t := range tm.tasks {
		if tm.tasks[i].done {
			tm.ListTask(t)
		}
	}
}

func (tm *TaskListManager) ListOverdueTasks() {
	for i, t := range tm.tasks {
		t1Date, _ := time.Parse("2006-01-02", tm.tasks[i].deadline)
		if t1Date.Before(time.Now()) && tm.tasks[i].done {
			tm.ListTask(t)
		}
	}
}

func (tm *TaskListManager) ListAll() {
	for _,t:= range tm.tasks {
		tm.ListTask(t)
	}
}

func main() {
	t1 := TaskList{0, "Hojiakbar", "design", "2019-12-01", false}
	t2 := TaskList{1, "Temur", "backend", "2019-12-05", true}
	t3 := TaskList{2, "Sardor", "frontend", "2019-10-31", false}
	t4 := TaskList{3, "Akbar", "backend", "2019-09-01", true}
	//t5 := TaskList{2, "Izzat", "Project manager", "2019-09-24", false}
	tm := NewTaskListManager()
	tm.Add(t1)
	tm.Add(t2)
	tm.Add(t3)
	tm.Add(t4)
	fmt.Println("List All")
	tm.ListAll()
	fmt.Println("List Unfinished Tasks")
	tm.ListUnfinishedTasks()
	tm.UpdateTitle(t2, "updated")
	tm.MakeDone(t2)
	fmt.Println("Make done")
	tm.ListUnfinishedTasks()
	fmt.Println("List overdue Tasks")
	tm.ListOverdueTasks()
	fmt.Println("Updated")
	tm.ListAll()
	fmt.Println("Deleted")
	tm.Delete(t1)
	tm.ListAll()
}
