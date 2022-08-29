package main

import (
	"log"
	"net/http"
	"verify/controller"
	"verify/dao"
)

func main() {
	log.Println("Starting Init DB.... ")
	dao.InitDB()
	log.Println("Starting Init handler...")
	err := http.ListenAndServe(":8080", controller.InitHandler())
	if err != nil {
		log.Panic("Init http server error:", err.Error())
	}
	log.Println("Congratulations servers init successfully")
}
