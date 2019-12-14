package models

type TaskList struct {
	Id       int    `db:"id"`
	Assignee string `db:"assignee"`
	Title    string `db:"title"`
	Deadline string `db:"deadline"`
	Done     bool   `db:"done"`
}


type TaskListInterface interface {
	Add(t TaskList) error
	UpdateTitle(id int, title string) error
	Delete(id int) error
	MakeDone(id int) error
	ListTask(id int) error
	ListUnfinishedTasks() error
	ListOverdueTasks() error
	ListAll() error
	GetAll() ([]TaskList,error)
}
