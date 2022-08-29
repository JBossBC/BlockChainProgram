package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	file, err := os.OpenFile("./routes.json", os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("open json file wrong :", err.Error())
		return
	}
	decoder := json.NewDecoder(file)
	var temp map[string]interface{}
	err = decoder.Decode(&temp)
	if err != nil {
		log.Fatal("decode json file wrong :", err.Error())
		return
	}
	for index, value := range temp {
		fmt.Printf("index: %s ,value:%s \n", index, value)
	}
}
