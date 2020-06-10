package api

import (
	"time"

	"sx/model"
	. "sx/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func LoginApi(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	password := Password2MD5(ctx.DefaultQuery("password", ""))

	if id == "" {
		ctx.JSON(200, gin.H{
			"err": 1,
			"msg": "用户ID不能为空.",
		})
		return
	}

	db := model.DB
	u := model.Person{
		Id:       id,
		Password: password,
	}
	r := db.First(&u)
	if r.RecordNotFound() {
		ctx.JSON(200, gin.H{
			"err":  2,
			"name": "无此用户或密码错误.",
		})
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    u.Id,
			"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
			"admin": u.Level == model.AdminLevel,
		})
		t, err := token.SignedString([]byte(JWTSECRET))
		if err == nil {
			//ctx.SetCookie("token", t, 60*60*24*7, "/", ctx.Request.Host, false, true)
			ctx.JSON(200, gin.H{
				"err":   0,
				"name":  u.Name,
				"level": u.Level,
				"admin": u.Level == model.AdminLevel,
				"token": t,
			})
		} else {
			//fmt.Println(err)
			ctx.JSON(200, gin.H{
				"err": 3,
				"msg": "登录失败.",
			})
		}
	}
}

func CheckLoginApi(ctx *gin.Context) {
	token := GetToken(ctx)
	if ok, id := GetId(token); ok {
		u := model.Person{
			Id: id,
		}
		if !model.DB.First(&u).RecordNotFound() {
			ctx.JSON(200, gin.H{
				"err":   0,
				"id":    u.Id,
				"name":  u.Name,
				"admin": model.AdminLevel == u.Level,
			})
			return
		}
	}
	ctx.JSON(200, gin.H{
		"err": 401,
		"msg": "未登录或登录过期.",
	})
}
