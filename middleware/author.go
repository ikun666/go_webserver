package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikun666/go_webserver/control"
	"github.com/ikun666/go_webserver/utils"
)

const (
	ACCESS_TOKEN    = "accessToken"
	REFRESH_TOKEN   = "refreshToken"
	ERR_TOKEN_NULL  = 40100
	ERR_TOKEN_PARSE = 40101
	ERR_TOKEN_GEN   = 40102
)

func Author() func(*gin.Context) {
	return func(ctx *gin.Context) {
		//从Header获取token
		accessToken := ctx.GetHeader(ACCESS_TOKEN)
		refreshToken := ctx.GetHeader(REFRESH_TOKEN)
		if accessToken == "" || refreshToken == "" {
			control.Fail(ctx, control.ResponseJson{
				Code: ERR_TOKEN_NULL,
				Msg:  "token is null",
			})
			return
		}
		claims, ok, err := utils.ParseToken(accessToken, refreshToken)
		if err != nil {
			control.Fail(ctx, control.ResponseJson{
				Code: ERR_TOKEN_PARSE,
				Msg:  "token prase err",
			})
			return
		}
		//需要刷新token
		if ok {
			token, err := utils.GetToken(claims.ID, claims.Name)
			if err != nil {
				control.Fail(ctx, control.ResponseJson{
					Code: ERR_TOKEN_GEN,
					Msg:  "token gen err",
				})
				return
			}
			//不用abort，后续操作不会中断
			ctx.JSON(http.StatusOK, control.ResponseJson{
				Msg:  "refresh token",
				Data: token,
			})
		}

		ctx.Next()

	}
}
