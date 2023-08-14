package control

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseJson struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Count  int64  `json:"count,omitempty"`
}

func OK(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(http.StatusOK, resp)
}
func Fail(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, resp)
}
func ServerFail(ctx *gin.Context, resp ResponseJson) {
	ctx.AbortWithStatusJSON(http.StatusInternalServerError, resp)
}
