package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikun666/go_webserver/control"
	"github.com/ikun666/go_webserver/global"
	"github.com/ikun666/go_webserver/utils"
)

const (
	TOKEN              = "Token"
	ERR_TOKEN_NULL     = 40100
	ERR_TOKEN_PARSE    = 40101
	ERR_TOKEN_UNEQUAL  = 40102
	ERR_TOKEN_EXPIRE   = 40103
	ERR_TOKEN_GENERATE = 40104
	ERR_TOKEN_REDIS    = 40105
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
		//如果解析失败
		claim, err := utils.ParseToken(token)
		if err != nil || claim.ID == 0 {
			control.Fail(ctx, control.ResponseJson{
				Msg:  "token prase err",
				Code: ERR_TOKEN_PARSE,
			})
			return
		}
		//解析token与登录返回token不一致
		redisKey := fmt.Sprintf("userLogin%d", claim.ID)
		redisToken, err := global.RedisClient.Get(redisKey)
		if err != nil || redisToken != token {
			control.Fail(ctx, control.ResponseJson{
				Status: http.StatusUnauthorized,
				Msg:    "token unequal",
				Code:   ERR_TOKEN_UNEQUAL,
			})
			return
		}
		//token过期
		expireDuration, err := global.RedisClient.GetExpireTime(redisKey)
		if err != nil || expireDuration <= 0 {
			control.Fail(ctx, control.ResponseJson{
				Status: http.StatusUnauthorized,
				Msg:    "token expire",
				Code:   ERR_TOKEN_EXPIRE,
			})
			return
		}
		//token有效期内访问续期

		//生成token
		token, err = utils.GenerateToken(claim.ID, claim.Name)
		if err != nil {
			control.Fail(ctx, control.ResponseJson{
				Status: http.StatusUnauthorized,
				Msg:    "generate token err",
				Code:   ERR_TOKEN_GENERATE,
			})
			return
		}
		//生成redis
		err = global.RedisClient.Set(redisKey, token)
		if err != nil {
			control.Fail(ctx, control.ResponseJson{
				Status: http.StatusUnauthorized,
				Msg:    "generate redis err",
				Code:   ERR_TOKEN_REDIS,
			})
			return
		}
		ctx.Header(TOKEN, token)

		ctx.Next()

	}
}
