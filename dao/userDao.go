package dao

import (
	"errors"

	"github.com/ikun666/go_webserver/dto"
	"github.com/ikun666/go_webserver/global"
	"github.com/ikun666/go_webserver/model"
	"github.com/ikun666/go_webserver/utils"
)

type UserDao struct {
	BaseDao
}

// 单例
var iUserDao *UserDao

func NewUserDao() *UserDao {
	if iUserDao == nil {
		return &UserDao{
			BaseDao: BaseDao{
				DB: global.DB,
			},
		}
	}
	return iUserDao
}

// 通过name获取用户
func (m *UserDao) GetUserByName(name string) (model.User, error) {
	var user model.User
	err := m.DB.Model(&user).Where("name=?", name).Find(&user).Error
	if err == nil && user.ID == 0 {
		err = errors.New("user not exist")
	}
	return user, err
}

// 添加用户
func (m *UserDao) AddUser(iAddUserDTO *dto.AddUserDTO) error {
	var user model.User
	iAddUserDTO.Convert2Model(&user)
	err := m.DB.Save(&user).Error
	//将数据库生成id回传并把密码置空用于回传客户
	if err == nil {
		iAddUserDTO.ID = user.ID
		iAddUserDTO.Password = ""
	}
	return err
}

// 登录
func (m *UserDao) Login(iLoginDTO *dto.LoginDTO) (model.User, utils.Tokens, error) {
	// var user model.User
	// err := m.DB.Model(&user).Where("name=? and password=?", iLoginDTO.Name, iLoginDTO.Password).Find(&user).Error
	user, err := m.GetUserByName(iLoginDTO.Name)
	//密码不对
	if err != nil || !utils.ComparePassword(user.Password, iLoginDTO.Password) {
		err = errors.New("password err")
		return user, utils.Tokens{}, err
	} else {
		token, err := utils.GetToken(user.ID, user.Name)

		if err != nil {
			err = errors.New("get token err")
		}
		return user, token, err
	}
}

// 删除用户
func (m *UserDao) DeleteUserByName(name string) error {
	var user model.User
	// 加上Unscoped() 物理删除 没有就是逻辑删除
	//如果 name 不存在 也会返回删除成功 并无影响
	//也可以先查询name是否存在返回err 但要多一次查询
	return m.DB.Unscoped().Model(&user).Where("name=?", name).Delete(&user).Error
}

// 更新用户
func (m *UserDao) UpdateUserByName(iUpdateUserDTO *dto.UpdateUserDTO) error {
	var user model.User
	user, err := m.GetUserByName(iUpdateUserDTO.Name)
	if err != nil {
		return err
	}
	iUpdateUserDTO.Convert2Model(&user)
	return m.DB.Save(&user).Error
}
