package service

import (
	"github.com/ikun666/go_webserver/dao"
	"github.com/ikun666/go_webserver/dto"
	"github.com/ikun666/go_webserver/model"
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

// 登录用户
func (s *UserService) Login(iLoginDTO *dto.LoginDTO) (model.User, string, error) {
	return s.Dao.Login(iLoginDTO)
}

// 查找用户
func (s *UserService) GetUserByName(name string) (model.User, error) {
	return s.Dao.GetUserByName(name)
}

// 删除用户
func (s *UserService) DeleteUserByName(name string) error {
	return s.Dao.DeleteUserByName(name)
}

// 更新用户
func (s *UserService) UpdateUserByName(iUpdateUserDTO *dto.UpdateUserDTO) error {
	return s.Dao.UpdateUserByName(iUpdateUserDTO)
}
