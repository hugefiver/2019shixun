package api

import (
	"io/ioutil"
	"strings"
	"sx/model"

	"github.com/gin-gonic/gin"
)

const MAX_FILE_SIZE = 8 * 1024 * 1024 * 4

var TypeMap = model.TypeMap

func GetImage(ctx *gin.Context) {
	p := model.Image{
		Id: ctx.Param("id"),
	}
	if model.DB.First(&p).RecordNotFound() {
		ctx.Status(404)
	} else {
		ctx.Data(200, p.Type, p.File)
	}
}

func UpdateImage(ctx *gin.Context) {
	id := ctx.Param("id")

	if file, err := ctx.FormFile("file"); err != nil {
		ctx.JSON(200, gin.H{
			"err": 1,
			"msg": "请求失败",
		})
	} else if file.Size > MAX_FILE_SIZE {
		ctx.JSON(200, gin.H{
			"err": 2,
			"msg": "文件过大",
		})
	} else {
		fs := strings.Split(file.Filename, ".")
		ext := fs[len(fs)-1]

		if t, ok := TypeMap[ext]; ok {
			f, _ := file.Open()
			bytes, _ := ioutil.ReadAll(f)
			p := model.Image{
				Id:   id,
				Type: t,
				File: bytes,
			}
			if model.DB.Where("id = ?", id).First(&model.Image{}).RecordNotFound() {
				model.DB.Create(&p)
			} else {
				model.DB.Model(&p).Where("id = ?", id).Update(p)
			}
			ctx.JSON(200, gin.H{
				"err": 0,
			})
		} else {
			ctx.JSON(200, gin.H{
				"err": 3,
				"msg": "不支持该类型文件.",
			})
		}
	}
}
