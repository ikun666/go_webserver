package control

import (
	"fmt"

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
	ERR_BIND         = 40000
	ERR_ADD_USER     = 40001
	ERR_LOGIN        = 40002
	ERR_GET_NAME     = 40003
	ERR_DELETE_NAME  = 40004
	ERR_UNPDATE_NAME = 40005
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
		return
	}

	user, token, err := c.Service.Login(&iLoginDTO)
	// fmt.Println(token)
	if err != nil {
		c.Fail(ctx, ResponseJson{
			Code: ERR_LOGIN,
			Msg:  err.Error(),
		})
		return
	}
	ctx.Set(token, user.Name)
	c.OK(ctx, ResponseJson{
		Msg: "login success",
		Data: gin.H{
			"user":  user,
			"token": token,
		},
	})
}

func (c UserControl) GetUserByName(ctx *gin.Context) {
	var iCommonDTO dto.CommonDTO
	err := ctx.ShouldBind(&iCommonDTO)
	if err != nil {
		c.Fail(ctx, ResponseJson{
			Code: ERR_BIND,
			Msg:  err.Error(),
		})
		return
	}
	user, err := c.Service.GetUserByName(iCommonDTO.Name)
	if err != nil {
		c.ServerFail(ctx, ResponseJson{
			Code: ERR_GET_NAME,
			Msg:  err.Error(),
		})
		return
	}
	c.OK(ctx, ResponseJson{
		Msg:  "get user success",
		Data: user,
	})
}

func (c UserControl) DeleteUserByName(ctx *gin.Context) {
	var iCommonDTO dto.CommonDTO
	err := ctx.ShouldBind(&iCommonDTO)
	if err != nil {
		c.Fail(ctx, ResponseJson{
			Code: ERR_BIND,
			Msg:  err.Error(),
		})
		return
	}
	err = c.Service.DeleteUserByName(iCommonDTO.Name)
	if err != nil {
		c.ServerFail(ctx, ResponseJson{
			Code: ERR_DELETE_NAME,
			Msg:  err.Error(),
		})
		return
	}
	c.OK(ctx, ResponseJson{
		Msg: "delete user success",
	})
}

func (c UserControl) UpdateUserByName(ctx *gin.Context) {
	var iUpdateUserDTO dto.UpdateUserDTO
	err := ctx.ShouldBind(&iUpdateUserDTO)
	if err != nil {
		c.Fail(ctx, ResponseJson{
			Code: ERR_BIND,
			Msg:  err.Error(),
		})
		return
	}
	err = c.Service.UpdateUserByName(&iUpdateUserDTO)
	if err != nil {
		c.ServerFail(ctx, ResponseJson{
			Code: ERR_UNPDATE_NAME,
			Msg:  err.Error(),
		})
		return
	}
	c.OK(ctx, ResponseJson{
		Msg: "update user success",
	})
}
func (c UserControl) TestApi(ctx *gin.Context) {
	var iCommonDTO dto.CommonDTO
	err := ctx.ShouldBind(&iCommonDTO)
	if err != nil {
		c.Fail(ctx, ResponseJson{
			Code: ERR_BIND,
			Msg:  err.Error(),
		})
		return
	}
	fmt.Println(iCommonDTO)
	c.OK(ctx, ResponseJson{
		Msg:  "test success",
		Data: iCommonDTO,
	})
}
