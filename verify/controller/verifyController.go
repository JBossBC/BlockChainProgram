package controller

import (
	"net/http"
	"verify/util"
)

type VerifyResponse struct {
	Result bool                   `json:result`
	Err    error                  `json:err`
	Data   map[string]interface{} `json:data`
}

func VerifyUser(w *http.ResponseWriter, r *http.Request) {
	var resultMap, err = util.ConvertRequestByJSON(r)
	var responseBody = &VerifyResponse{}
	if err != nil {
		responseBody.Err = err
		responseBody.Result = false
	}
	resultMap["name"]

}
