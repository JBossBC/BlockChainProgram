package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"verify/service"
	"verify/util"
)

type VerifyResponse struct {
	Result bool  `json:result`
	Err    error `json:err`
}

func VerifyUser(w *http.ResponseWriter, r *http.Request) {
	var resultMap, err = util.ConvertRequestByJSON(r)
	var responseBody = &VerifyResponse{}
	if err != nil {
		responseBody.Err = err
		responseBody.Result = false
		bytes, _ := json.Marshal(responseBody)
		io.WriteString(*w, string(bytes))
	}
	result := service.GetVeiryService().VerifyUser(resultMap)
	if result == util.Success {
		responseBody.Result = true
		responseBody.Err = nil
	} else {
		responseBody.Result = false
		responseBody.Err = fmt.Errorf("用户名或密码错误")
	}
	bytes, _ := json.Marshal(responseBody)
	io.WriteString(*w, string(bytes))
}
