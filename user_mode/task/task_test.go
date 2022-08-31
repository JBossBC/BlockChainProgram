package task

import (
	"fmt"
	"testing"
)

func TestTask(test *testing.T) {
	task := &AbstractTask{}
	task.Start()
	for s, taskInterface := range task.tasks {
		fmt.Println(s)
		fmt.Println(taskInterface)
	}
}
