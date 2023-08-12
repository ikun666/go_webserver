package dao

import (
	"github.com/ikun666/go_webserver/dto"
	"github.com/ikun666/go_webserver/global"
	"github.com/ikun666/go_webserver/model"
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
