package control

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseControl struct {
	// Ctx *gin.Context
}

// type BuildRequestOption struct {
// 	Ctx *gin.Context
// 	DTO any
// }

// func (c BaseControl) BuildRequest(opt BuildRequestOption) error {
// 	//绑定上下文以及DTO
// 	c.Ctx = opt.Ctx
// 	err := c.Ctx.ShouldBind(opt.DTO)
// 	return err
// }

func (c BaseControl) OK(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}
func (c BaseControl) Fail(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
}
func (c BaseControl) ServerFail(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}
