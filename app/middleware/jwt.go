package middleware

import (
	"Raven-gin/app/common/response"
	g "Raven-gin/global"
	"Raven-gin/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.FailByAuth(ctx, "没有访问权限", 0)
			ctx.Abort()
			return
		}
		tokenStr = tokenStr[len(g.Cof.Jwt.TokenType)+1:]

		token, err := jwt.ParseWithClaims(tokenStr, &utils.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(g.Cof.Jwt.Secret), nil
		})

		if err != nil || utils.IsInBlacklist(tokenStr) {
			response.FailByAuth(ctx, "token无效", 0)
			ctx.Abort()
			return
		}

		claims := token.Claims.(*utils.CustomClaims)

		if claims.Issuer != GuardName {
			response.FailByAuth(ctx, "Token签发者不正确", 0)
			ctx.Abort()
			return
		}

		ctx.Set("token", token)
		ctx.Set("id", claims.Id)
	}
}
