package dto

import "github.com/ikun666/go_webserver/model"

type AddUserDTO struct {
	ID       uint
	Name     string `json:"name" form:"name" binding:"required" message:"name is null"`
	RealName string `json:"realName" form:"realName"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password" binding:"required" message:"password is null"`
}

func (m *AddUserDTO) Convert2Model(iUser *model.User) {
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Avatar = m.Avatar
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
	iUser.Password = m.Password
}
