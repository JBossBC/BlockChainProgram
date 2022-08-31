package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"servers_distribute/util"
)

type myHandler struct {
}

type proxyResponse struct {
	AbstractRepsonse
}

var routes map[string]interface{}

//后期进行递归操作的优化(基于inode文件系统的操作)
func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := util.HandleURLToPath(r.URL.String())
	proxy := &proxyResponse{}
	switch context[0] {
	case "user_mode":
		//分发服务
		request, err := http.NewRequest(http.MethodPost, "http://localhost:8081/verify", r.Body)
		if err != nil {
			proxy.Message = err.Error()
			marshal, _ := json.Marshal(&proxy)
			w.Write(marshal)
			return
		}
		request.Header.Add("content-type", "application/json")
		client := http.Client{}
		do, err := client.Do(request)
		if err != nil {
			proxy.Message = err.Error()
			marshal, _ := json.Marshal(&proxy)
			w.Write(marshal)
			return
		}
		all, err := io.ReadAll(do.Body)
		if err != nil {
			proxy.Message = err.Error()
			marshal, _ := json.Marshal(&proxy)
			w.Write(marshal)
			return
		}
		w.Write(all)
	}
}
func GetHandler() *myHandler {
	return &myHandler{}
}

//TODO 增加可扩展性
//func MatchingContext(context []string) {
//	var result bool
//	for i := 0; i < len(context); i++ {
//
//	}
//}
//TODO error :jsonfile hasn't find
func InitRoutes() {
	log.Println("Starting init routes...")
	jsonFile, err := os.OpenFile("./Controller/routes.json", os.O_APPEND, 0644)
	if err != nil {
		log.Panic("routesFile init err :", err.Error())
		panic(fmt.Sprintf("routesFile init err :%s", err.Error()))
	}
	err = json.NewDecoder(jsonFile).Decode(&routes)
	if err != nil {
		log.Panic("decoder routes file err :", err.Error())
		panic(fmt.Sprintf("routesFile init err :%s", err.Error()))
	}
	log.Println("routes file init success")
}
