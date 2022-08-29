package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"user_mode/service"
	"user_mode/util"
)

type VerifyResponse struct {
	AbstractRepsonse
}

func VerifyUser(w *http.ResponseWriter, r *http.Request) {
	var resultMap, err = util.ConvertRequestByJSON(r)
	var responseBody = &VerifyResponse{}
	if err != nil {
		responseBody.Message = err.Error()
		responseBody.Result = false
		bytes, _ := json.Marshal(responseBody)
		io.WriteString(*w, string(bytes))
	}
	result := service.GetVeiryService().VerifyUser(resultMap)
	if result == util.Success {
		responseBody.Result = true
		//无信息
		responseBody.Message = ""
	} else {
		responseBody.Result = false
		responseBody.Message = fmt.Errorf("用户名或密码错误").Error()
	}
	bytes, _ := json.Marshal(responseBody)
	io.WriteString(*w, string(bytes))
}
