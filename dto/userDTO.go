package dto

import "github.com/ikun666/go_webserver/model"

type AddUserDTO struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	RealName string `json:"realName" form:"realName"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password,omitempty" form:"password" binding:"required"`
}

func (m *AddUserDTO) Convert2Model(iUser *model.User) {
	iUser.Name = m.Name
	iUser.RealName = m.RealName
	iUser.Avatar = m.Avatar
	iUser.Mobile = m.Mobile
	iUser.Email = m.Email
	iUser.Password = m.Password
}

type LoginDTO struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `json:"password,omitempty" form:"password" binding:"required"`
}

type CommonDTO struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

type UpdateUserDTO struct {
	Name     string `json:"name" form:"name"`
	RealName string `json:"realName" form:"realName"`
	Avatar   string `json:"avatar" form:"avatar"`
	Mobile   string `json:"mobile" form:"mobile"`
	Email    string `json:"email" form:"email"`
}

func (m *UpdateUserDTO) Convert2Model(iUser *model.User) {
	if m.RealName != "" {
		iUser.RealName = m.RealName
	}
	if m.Avatar != "" {
		iUser.Avatar = m.Avatar
	}
	if m.Mobile != "" {
		iUser.Mobile = m.Mobile
	}
	if m.Email != "" {
		iUser.Email = m.Email
	}

}
