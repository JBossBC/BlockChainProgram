package service

import (
	"sync"
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
func (v *VerifyService) verifyUser(information map[string]interface{}) util.StatusCode {

}
