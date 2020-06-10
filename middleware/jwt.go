package middleware

import (
	"sx/model"
	. "sx/util"

	"github.com/gin-gonic/gin"
)

func LoginOnlyMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := GetToken(ctx)
		if ok, _ := GetId(token); ok {
			ctx.Next()
			return
		}
		ctx.JSON(200, gin.H{
			"err": 401,
			"msg": "未登录.",
		})
		ctx.Abort()
	}
}

func AdminOnlyMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := GetToken(ctx)
		ok, id := GetId(token)
		if !ok {
			ctx.JSON(200, gin.H{
				"err": 401,
				"msg": "未登录.",
			})
			ctx.Abort()
		} else {
			u := model.Person{
				Id: id,
			}
			if model.DB.First(&u).RecordNotFound() {
				ctx.JSON(200, gin.H{
					"err": 401,
					"msg": "未登录.",
				})
				ctx.Abort()
			} else if u.Level != model.AdminLevel {
				ctx.JSON(200, gin.H{
					"err": 403,
					"msg": "权限不足.",
				})
				ctx.Abort()
			} else {
				ctx.Next()
			}
		}
	}
}

func NotLoginRedirectMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := GetToken(ctx)
		if ok, _ := GetId(token); ok {
			ctx.Next()
			return
		}
		ctx.Redirect(302, "/login")
		ctx.Abort()
	}
}
