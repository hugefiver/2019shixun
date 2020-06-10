package api

import (
	"sx/model"

	"github.com/gin-gonic/gin"
)

func UpdateMochine(ctx *gin.Context) {
	id := ctx.Param("id")
	m := model.Mochine{
		Id: id,
	}

	if err := ctx.ShouldBindJSON(&m); err == nil {
		m.Id = id
		if m.LifeYear <= 0 || m.Kind == "" || m.AdminId == "" {
			ctx.JSON(200, gin.H{
				"err": 2,
				"msg": "数据不正确",
			})
		} else {
			if model.DB.Where("id = ?", id).First(&model.Mochine{}).RecordNotFound() {
				model.DB.Create(&m)
				ctx.JSON(200, gin.H{
					"err":    0,
					"status": "create",
				})
			} else {
				model.DB.Where("id = ?", id).First(&model.Mochine{}).Update(m)
				ctx.JSON(200, gin.H{
					"err":    0,
					"status": "update",
				})
			}
		}
	} else {
		ctx.JSON(200, gin.H{
			"err": 1,
			"msg": "请求错误",
		})
	}
}

func GetMochineInfo(ctx *gin.Context) {
	id := ctx.Param("id")
	m := model.Mochine{}

	if model.DB.Where("id = ?", id).First(&m).RecordNotFound() {
		ctx.JSON(404, gin.H{
			"err": 2,
		})
	} else {
		ctx.JSON(200, gin.H{
			"err":   0,
			"id":    m.Id,
			"admin": m.AdminId,
			"kind":  m.Kind,
			"life":  m.LifeYear,
		})
	}
}

func DeleteMochine(ctx *gin.Context) {
	id := ctx.Param("id")

	if model.DB.Where("id = ?", id).First(&model.Mochine{}).RecordNotFound() {
		ctx.JSON(200, gin.H{
			"err": 2,
			"msg": "Not Found.",
		})
	} else {
		model.DB.Where("id = ?", id).Delete(&model.Mochine{})
		ctx.JSON(200, gin.H{
			"err": 0,
		})
	}
}

func GetMochines(ctx *gin.Context) {
	var ps []model.Mochine

	q := model.DB
	if kind := ctx.Query("kind"); kind != "" {
		q = q.Where("kind = ?", kind)
	}
	q.Find(&ps)

	list := make([]string, len(ps))
	for i, p := range ps {
		list[i] = p.Id
	}
	ctx.JSON(200, gin.H{
		"err":      0,
		"mochines": ps,
	})
}
