package api

import (
	"sx/model"
	. "sx/util"
	"time"

	"github.com/gin-gonic/gin"
)

func UseMochine(c *gin.Context) {
	_, user := GetId(GetToken(c))
	id := c.Param("id")

	model.DB.Create(&model.McUse{
		UserId:    user,
		MochineId: id,
		UseTime:   time.Now(),
	})

	c.JSON(200, gin.H{
		"err": 0,
	})
}

func CheckUse(c *gin.Context) {
	id := c.Param("id")
	var us []model.McUse

	q := model.DB.Where("mochine_id = ?", id)
	q.Find(&us)

	c.JSON(200, gin.H{
		"err":  0,
		"uses": us,
	})
}
