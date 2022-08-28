package controller

import (
	"net/http"
	"strings"
)

type MyHandler struct {
}

func InitHandler() *MyHandler {
	return &MyHandler{}
}

//v1.0扩展性可能不够,不能成为框架
func (handler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlStr := r.URL.String()
	pathArr := handleURL(urlStr)
	switch pathArr[0] {
	case "verify":
		VerifyUser(&w, r)
	}
}

func handleURL(url string) []string {
	//使用rune防止因为编码集导致出错
	var result = len(url) - 1
	for index, value := range url {
		if value == '?' {
			result = index - 1
			break
		}
	}
	var discussRawQuery = url[:result+1]
	return strings.Split(discussRawQuery, "/")[1:]
}
