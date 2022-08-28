package main

import (
	"net/http"
	"verify/controller"
	"verify/dao"
)

func main() {
	dao.InitDB()
	http.ListenAndServe(":8080", controller.InitHandler())
}
