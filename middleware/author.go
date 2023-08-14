package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/ikun666/go_webserver/control"
	"github.com/ikun666/go_webserver/utils"
)

const (
	TOKEN           = "token"
	ERR_TOKEN_NULL  = 40100
	ERR_TOKEN_PARSE = 40101
)

func Author() func(*gin.Context) {
	return func(ctx *gin.Context) {
		//从Header获取token
		token := ctx.GetHeader(TOKEN)
		if token == "" {
			control.Fail(ctx, control.ResponseJson{
				Code: ERR_TOKEN_NULL,
				Msg:  "token is null",
			})
			return
		}
		claim, err := utils.ParseToken(token)
		if err != nil || claim.ID == 0 {
			control.Fail(ctx, control.ResponseJson{
				Code: ERR_TOKEN_PARSE,
				Msg:  "token prase err",
			})
			return
		}
		// control.OK(ctx, control.ResponseJson{
		// 	Msg:  "author",
		// 	Data: claim,
		// })
		ctx.Next()

	}
}
