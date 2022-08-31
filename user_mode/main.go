package main

import (
	"log"
	"net/http"
	"user_mode/controller"
	"user_mode/dao"
	"user_mode/task"
)

func main() {
	dao.InitDB()
	task.StartTask()
	err := http.ListenAndServe(":8080", controller.InitHandler())
	if err != nil {
		log.Panic("Init http server error:", err.Error())
	}
	log.Println("Congratulations servers init successfully")
}
