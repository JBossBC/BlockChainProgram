package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//https://localhost:8080/  ? username=?&password=?&
//是直接解析为对象还是解析为map???? V1.0解析为map可迭代性可能会更高
func ConvertRequestByJSON(r *http.Request) (map[string]interface{}, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("Json convert error:%s", err.Error())
	}
	var result = make(map[string]interface{})
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return nil, fmt.Errorf("json convert error: %s", err.Error())
	}
	return result, nil
}

//系统封装,可以不用解析错误
func ConvertJsonByObj(obj interface{}) []byte {
	var result, err = json.Marshal(&obj)
	if err != nil {
		return nil
	}
	return result
}
