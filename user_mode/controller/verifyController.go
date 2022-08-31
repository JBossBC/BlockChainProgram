package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"user_mode/service"
	"user_mode/util"
)

func VerifyUser(w *http.ResponseWriter, r *http.Request) {
	var resultMap, err = util.ConvertRequestByJSON(r)
	var responseBody = &AbstractRepsonse{}
	defer handleDataToResponse(w, responseBody)
	if err != nil {
		responseBody.Message = err.Error()
		return
	}
	result := service.GetVeiryService().VerifyUser(resultMap)
	if result == util.Success {
		responseBody.Result = true
		//无信息
		responseBody.Message = ""
	} else {
		responseBody.Message = fmt.Errorf("用户名或密码错误").Error()
	}
}
func CreateUser(w *http.ResponseWriter, r *http.Request) {
	data, err := util.ConvertRequestByJSON(r)
	var responseBody = &AbstractRepsonse{}
	defer handleDataToResponse(w, responseBody)
	if data["username"] == nil || data["password"] == nil {
		responseBody.Message = fmt.Sprintf("information defect")
		return
	}
	if err != nil {
		responseBody.Message = err.Error()
		return
	}
	statusCode := service.GetVeiryService().CreateUser(data)
	if statusCode != util.Success {
		responseBody.Message = "create user failed"
	} else {
		responseBody.Result = true
		responseBody.Message = "create user success"
		log.Printf("create user {username:%s,password:%s} ", data["username"], data["password"])
	}
}
func UpdateUser(w *http.ResponseWriter, r *http.Request) {
	data, err := util.ConvertRequestByJSON(r)
	var responseBody = &AbstractRepsonse{}
	defer handleDataToResponse(w, responseBody)
	if data["username"] == nil || data["password"] == nil {
		responseBody.Message = fmt.Sprintf("information defect")
		return
	}
	if err != nil {
		responseBody.Message = err.Error()
		return
	}
	statusCode := service.GetVeiryService().UpdateUserInfo(data)
	if statusCode != util.Success {
		responseBody.Message = "update user failed"
	} else {
		responseBody.Result = true
		responseBody.Message = "update user success"
	}
}
func DeleteUser(w *http.ResponseWriter, r *http.Request) {
	data, err := util.ConvertRequestByJSON(r)
	var responseBody = &AbstractRepsonse{}
	defer handleDataToResponse(w, responseBody)
	if data["username"] == nil || data["password"] == nil {
		responseBody.Message = fmt.Sprintf("information defect")
		return
	}
	if err != nil {
		responseBody.Message = err.Error()
		return
	}
	statusCode := service.GetVeiryService().UpdateUserInfo(data)
	if statusCode != util.Success {
		responseBody.Message = "delete user failed"
	} else {
		responseBody.Result = true
		responseBody.Message = "delete user success"
	}
}
func handleDataToResponse(w *http.ResponseWriter, repsonse *AbstractRepsonse) {
	bytes, _ := json.Marshal(repsonse)
	io.WriteString(*w, string(bytes))
}
