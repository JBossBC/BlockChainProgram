package controller

import (
	"log"
	"net/http"
	"user_mode/util"
)

type MyHandler struct {
}

func InitHandler() *MyHandler {
	log.Println("Starting server handler...")
	return &MyHandler{}
}

//v1.0扩展性可能不够,不能成为框架
func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlStr := r.URL.String()
	pathArr := util.HandleURLToPath(urlStr)
	switch pathArr[0] {
	case "verify":
		VerifyUser(&w, r)
	case "create":
		CreateUser(&w, r)
	case "update":
		UpdateUser(&w, r)
	case "delete":
		DeleteUser(&w, r)
	}
}
