package control

import (
	"net/http"

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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseJson{
			Code: ERR_BIND,
			Msg:  err.Error(),
		})
		return
	}
	err = c.Service.AddUser(&iAddUserDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ResponseJson{
			Code: ERR_ADD_USER,
			Msg:  err.Error(),
		})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, ResponseJson{
		Msg:  "add user success",
		Data: iAddUserDTO,
	})
}
