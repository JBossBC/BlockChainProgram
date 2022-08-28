package util

import "fmt"

func InterfaceConvertString(data interface{}) (string, error) {
	switch data.(type) {
	case string:
		return data.(string), nil
	default:
		return "", fmt.Errorf("this data convert to string error")
	}
}
