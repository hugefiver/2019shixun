package util

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const JWTSECRET = "alpha"

func GetToken(ctx *gin.Context) (token string) {
	token = ctx.GetHeader("token")
	if token == "" {
		var err error
		token, err = ctx.Cookie("token")
		if err != nil {
			token = ""
		}
	}
	return
}

func GetId(token string) (ok bool, id string) {
	if token == "" {
		return false, ""
	} else {
		t, err := jwt.Parse(token, func(token *jwt.Token) (i interface{}, e error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(JWTSECRET), nil
		})
		if err != nil {
			return false, ""
		} else if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid &&
			int64(claims["exp"].(float64)) > time.Now().Unix() {
			return ok, claims["id"].(string)
		} else {
			return false, ""
		}
	}
}
