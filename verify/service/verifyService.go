package service

import (
	"strings"
	"sync"
	"verify/dao"
	"verify/util"
)

//单例模式

type VerifyService struct {
}

var (
	verify *VerifyService
	once   sync.Once
)

//sync.Once可以防止DCL
func GetVeiryService() *VerifyService {
	//内联优化
	once.Do(func() {
		verify = &VerifyService{}
	})
	return verify
}
func (v *VerifyService) VerifyUser(information map[string]interface{}) util.StatusCode {
	userDao := dao.GetUserDao()
	var usernameStr, err = util.InterfaceConvertString(information["usernmame"])
	if err != nil {
		return util.UserNameError
	}
	passwordStr, err := util.InterfaceConvertString(information["password"])
	cryptoPassword := util.MD5Crypto(passwordStr)
	info := userDao.GetUserInfo(usernameStr)
	if strings.Compare(cryptoPassword, info.Password) == 0 {
		return util.Success
	}
	return util.PasswordError
}
