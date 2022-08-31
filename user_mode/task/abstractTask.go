package task

import (
	"log"
	"reflect"
)

var TaskSet *AbstractTask

func StartTask() {
	log.Println("Starting crontab task...")
	TaskSet = &AbstractTask{}
	TaskSet.Start()
	log.Println("task is running success")
}

//组合模式
type AbstractTask struct {
	tasks map[string]TaskInterface
}

func (a *AbstractTask) Start() {
	a.tasks = make(map[string]TaskInterface, 1000)
	a.addTask()
	for _, value := range a.tasks {
		value.Start()
	}
}

//添加任务
func (a *AbstractTask) addTask() {
	var deleteDaoTask = &DeleteUserTask{}
	a.tasks[reflect.TypeOf(*deleteDaoTask).String()] = deleteDaoTask
}

type TaskInterface interface {
	Start()
}
