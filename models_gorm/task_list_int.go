package models

import "github.com/jinzhu/gorm"

type TaskList struct {
	gorm.Model
	Assignee string
	Title    string
	Deadline string
	Done     bool
}


type TaskListInterface interface {
	Add(t TaskList) error
	UpdateTitle(id uint, title string) error
	Delete(id uint) error
	MakeDone(id uint) error
	ListTask(id uint) error
	ListUnfinishedTasks() error
	ListOverdueTasks() error
	ListAll() error
	GetAll() ([]TaskList,error)
}
