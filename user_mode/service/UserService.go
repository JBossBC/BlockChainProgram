package service

import (
	"strings"
	"sync"
	"time"
	"user_mode/dao"
	"user_mode/util"
)

//单例模式

type UserService struct {
}

var (
	userService *UserService
	once        sync.Once
)

//sync.Once可以防止DCL
func GetVeiryService() *UserService {
	//内联优化
	once.Do(func() {
		userService = &UserService{}
	})
	return userService
}
func (u *UserService) VerifyUser(information map[string]interface{}) util.StatusCode {
	userDao := dao.GetUserDao()
	user := u.handleMapToUser(information)
	cryptoPassword := util.MD5Crypto(user.Password)
	info, err := userDao.GetUserInfo(user.Username)
	if err != nil {
		return util.SystemError
	}
	if strings.Compare(cryptoPassword, info.Password) == 0 {
		return util.Success
	}
	return util.PasswordError
}
func (u *UserService) CreateUser(information map[string]interface{}) util.StatusCode {
	user := u.handleMapToUser(information)

	err := dao.GetUserDao().InsertUserInfo(user)
	if err != nil {
		return util.SystemError
	}
	return util.Success
}

func (u *UserService) UpdateUserInfo(information map[string]interface{}) util.StatusCode {
	user := u.handleMapToUser(information)
	err := dao.GetUserDao().UpdateUserInfo(user)
	if err != nil {
		return util.SystemError
	}
	return util.Success
}
func (u *UserService) DeleteUserInfo(information map[string]interface{}) util.StatusCode {
	user := u.handleMapToUser(information)
	err := dao.GetUserDao().DeleteUserDao(user)
	if err != nil {
		return util.SystemError
	}
	return util.Success
}

func (u *UserService) handleMapToUser(information map[string]interface{}) *dao.User {
	username := util.InterfaceConvertString(information["username"])
	password := util.InterfaceConvertString(information["password"])
	cryptoPassword := util.MD5Crypto(password)
	user := &dao.User{
		Username:  username,
		Password:  cryptoPassword,
		Create_at: time.Now(),
		Update_at: time.Now(),
		Delete_at: time.Now(),
	}
	return user
}
