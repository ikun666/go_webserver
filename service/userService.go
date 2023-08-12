package service

import (
	"github.com/ikun666/go_webserver/dao"
	"github.com/ikun666/go_webserver/dto"
)

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

// 单例
var iUserService *UserService

func NewUserService() *UserService {
	if iUserService == nil {
		iUserService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return iUserService
}

// 添加用户
func (s *UserService) AddUser(iAddUserDTO *dto.AddUserDTO) error {
	return s.Dao.AddUser(iAddUserDTO)
}
