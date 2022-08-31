package task

import (
	"time"
	"user_mode/dao"
)

type DeleteUserTask struct {
}

func (*DeleteUserTask) Start() {
	time.AfterFunc(24*time.Hour, dao.GetUserDao().DeleteHardUser)
}
