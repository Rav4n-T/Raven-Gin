package middleware

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"Raven-gin/app/common/response"
	"Raven-gin/app/services"
	g "Raven-gin/global"
	"Raven-gin/utils"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenStr := ctx.Request.Header.Get("Authorization")
		if tokenStr == "" || len(tokenStr) <= len(g.Cof.Jwt.TokenType)+1 {
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

		if claims.ExpiresAt-time.Now().Unix() < g.Cof.Jwt.RefreshGracePeriod {
			lock := utils.Lock("reffresh_token_lock", g.Cof.Jwt.JwtBlacklistGracePeriod)
			if lock.Get() {
				err, user := services.UserService.GetUserInfoWhithGuardName(GuardName, claims.Id)
				if err != nil {
					g.Log.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := utils.CreateToken(GuardName, user)
					ctx.Header("new-token", tokenData.AccessToken)
					ctx.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
					_ = utils.JoinBlacklist(token)
					// 更新 ctx 中的 token ？
				}
			}
		}

		ctx.Set("token", token)
		ctx.Set("id", claims.Id)
	}
}
