package utils

import (
	g "Raven-gin/global"
	"time"

	"github.com/dgrijalva/jwt-go"
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
