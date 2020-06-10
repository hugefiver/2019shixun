package api

import (
	"sx/model"
	"sx/util"

	"github.com/gin-gonic/gin"
)

func GetPeople(ctx *gin.Context) {
	var ps []model.Person

	q := model.DB
	if name := ctx.Query("name"); name != "" {
		q = q.Where("name = ?", name)
	}
	if level := ctx.Query("level"); level != "" {
		q = q.Where("level = ?", level)
	}
	q.Find(&ps)

	list := make([]string, len(ps))
	for i, p := range ps {
		list[i] = p.Id
	}
	ctx.JSON(200, gin.H{
		"err":    0,
		"people": list,
	})
}

func AddPerson(ctx *gin.Context) {
	var p model.Person
	if ctx.ShouldBindJSON(&p) == nil {
		if model.DB.Where("id = ?", p.Id).First(&model.Person{}).RecordNotFound() {
			p.Password = util.Password2MD5(p.Password)
			model.DB.Create(&p)
			ctx.JSON(200, gin.H{
				"err": 0,
				"msg": "Success.",
			})
		} else {
			ctx.JSON(200, gin.H{
				"err": 2,
				"msg": "人员已存在",
			})
		}
	} else {
		ctx.JSON(200, gin.H{
			"err": 1,
			"msg": "请求错误.",
		})
	}
}

func GetPersonInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	p := model.Person{Id: id}

	if model.DB.First(&p).RecordNotFound() {
		ctx.JSON(404, gin.H{
			"err": 1,
			"msg": "用户不存在",
		})
	} else {
		ctx.JSON(200, gin.H{
			"err":   0,
			"id":    p.Id,
			"name":  p.Name,
			"level": p.Level,
		})
	}
}

func UpdatePersonInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	p := model.Person{}

	if ctx.ShouldBindJSON(&p) == nil {
		if p.Name == "" {
			ctx.JSON(200, gin.H{
				"err": 2,
				"msg": "用户名不能为空.",
			})
			return
		}
		var t model.Person
		if model.DB.First(&t, "id = ?", id).RecordNotFound() {
			p.Password = util.Password2MD5(p.Password)
			model.DB.Create(&p)
			ctx.JSON(200, gin.H{
				"err":    0,
				"status": "create",
			})
		} else {
			if p.Password == "" {
				p.Password = t.Password
			} else {
				p.Password = util.Password2MD5(p.Password)
			}
			model.DB.Model(&p).Where("id = ?", id).Updates(p)
			ctx.JSON(200, gin.H{
				"err":    0,
				"status": "update",
			})
		}
	} else {
		ctx.JSON(200, gin.H{
			"err": 1,
			"msg": "请求错误",
		})
	}
}

func DeletePerson(ctx *gin.Context) {
	id := ctx.Param("id")
	if model.DB.Where("id = ?", id).First(&model.Person{}).RecordNotFound() {
		ctx.JSON(200, gin.H{
			"err": 1,
			"msg": "用户不存在",
		})
	} else {
		model.DB.Where("id = ?", id).Delete(&model.Person{})
		ctx.JSON(200, gin.H{
			"err": 0,
		})
	}
}
