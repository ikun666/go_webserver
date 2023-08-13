package control

import (
	"github.com/gin-gonic/gin"
	"github.com/ikun666/go_webserver/dto"
	"github.com/ikun666/go_webserver/service"
)

type UserControl struct {
	BaseControl
	Service *service.UserService
}

// 返回错误码
const (
	ERR_BIND     = 40000
	ERR_ADD_USER = 40001
	ERR_LOGIN    = 40002
)

// 多用户访问，不做单例
func NewUserControl() UserControl {
	return UserControl{
		Service: service.NewUserService(),
	}
}

func (c UserControl) AddUser(ctx *gin.Context) {
	var iAddUserDTO dto.AddUserDTO
	// err := c.BuildRequest(BuildRequestOption{
	// 	Ctx: ctx,
	// 	DTO: &iAddUserDTO,
	// })

	err := ctx.ShouldBind(&iAddUserDTO)
	if err != nil {
		c.Fail(ctx, ResponseJson{
			Code: ERR_BIND,
			Msg:  err.Error(),
		})
		return
	}
	err = c.Service.AddUser(&iAddUserDTO)
	if err != nil {
		c.ServerFail(ctx, ResponseJson{
			Code: ERR_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	c.OK(ctx, ResponseJson{
		Msg:  "add user success",
		Data: iAddUserDTO,
	})
}

func (c UserControl) Login(ctx *gin.Context) {
	var iLoginDTO dto.LoginDTO
	err := ctx.ShouldBind(&iLoginDTO)
	if err != nil {
		c.Fail(ctx, ResponseJson{
			Code: ERR_BIND,
			Msg:  err.Error(),
		})
	}

	user, err := c.Service.Login(&iLoginDTO)
	if err != nil {
		c.ServerFail(ctx, ResponseJson{
			Code: ERR_LOGIN,
			Msg:  err.Error(),
		})
		return
	}
	c.OK(ctx, ResponseJson{
		Msg:  "login success",
		Data: user,
	})
}
