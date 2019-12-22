package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type GormDB struct{
	db *gorm.DB
}


func  NewTaskList() (TaskListInterface,error){
	nw := GormDB{}
	psqlInfo := fmt.Sprintf(" user=%s dbname=%s password=%s",username,dbname2,password)
	nw.db,err = gorm.Open("postgres",psqlInfo)
	nw.db.AutoMigrate(&TaskList{})
	if err!=nil {
		panic(err.Error())
		return nil,err
	}
	nw.db.DropTableIfExists(&TaskList{})
	nw.db.CreateTable(&TaskList{})
	return &nw,nil
}

func (g *GormDB) TableName() string {
	return "task_list_info"
}

func (g *GormDB) Add(t TaskList) error {
	err = g.db.Create(&t).Error
	return Error(err)
}

func (g *GormDB) UpdateTitle(id uint,title string) error {
	err = g.db.Model(&TaskList{}).Where("id = ?",id).Update("title",title).Error
	return Error(err)
}

func (g *GormDB) Delete(id uint) error{
	err = g.db.Where("id = ?",id).Delete(&TaskList{}).Error
	return Error(err)
}

func(g *GormDB) MakeDone(id uint) error {
	err = g.db.Model(&TaskList{}).Where("id = ?",id).Update("done = true").Error
	return Error(err)
}

func (g *GormDB) ListTask(id uint) error {
	var taskList []TaskList
	err = g.db.Where("id = ?",id).First(&taskList).Error
	return ErrorPrint(err,taskList)
}

func (g *GormDB) ListUnfinishedTasks() error {
	var taskList []TaskList
	err = g.db.Where("done = false ").Find(&taskList).Error
	return ErrorPrint(err,taskList)
}

func (g *GormDB) ListOverdueTasks() error {
	var taskList [] TaskList
	err = g.db.Where("done = false and deadline > now()").Find(&taskList).Error
	return ErrorPrint(err,taskList)
}

func (g *GormDB) ListAll() error {
	var taskList [] TaskList
	err = g.db.Find(&taskList).Error
	return ErrorPrint(err,taskList)
}

func (g *GormDB) GetAll() ([]TaskList, error) {
	var taskList [] TaskList
	err = g.db.Find(&taskList).Error
	return taskList,Error(err)
}

func Error(err error) error{
	if err!=nil {
		return err
	}
	return nil
}

func ErrorPrint(err error,taskList []TaskList) error {
	if err!=nil {
		return err
	}
	fmt.Println(taskList)
	return nil
}



