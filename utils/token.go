package utils

import (
	"context"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	g "Raven-Admin/global"
)

type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.StandardClaims
}

type TokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func CreateToken(GuardName string, user JwtUser) (tokenData TokenData, err error, token *jwt.Token) {
	token = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		CustomClaims{
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + g.Cof.Jwt.JwtExp,
				Id:        user.GetUid(),
				Issuer:    GuardName,
				NotBefore: time.Now().Unix() - 1000,
			},
		},
	)

	tokenStr, err := token.SignedString([]byte(g.Cof.Jwt.Secret))
	tokenData = TokenData{
		tokenStr,
		int(g.Cof.Jwt.JwtExp),
		g.Cof.Jwt.TokenType,
	}
	return
}

func getBlacklistKey(tokenStr string) string {
	return "jwt_blacklist:" + MD5([]byte(tokenStr))
}

func JoinBlacklist(token *jwt.Token) (err error) {
	nowUnix := time.Now().Unix()
	timer := time.Duration(token.Claims.(*CustomClaims).ExpiresAt-nowUnix) * time.Second
	err = g.Redis.SetNX(context.Background(), getBlacklistKey(token.Raw), nowUnix, timer).Err()
	return
}

func IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := g.Redis.Get(context.Background(), getBlacklistKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	if time.Now().Unix()-joinUnix < g.Cof.Jwt.JwtBlacklistGracePeriod {
		return false
	}
	return true
}
