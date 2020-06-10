package api

import (
	"sx/model"
	. "sx/util"
	"time"

	"github.com/gin-gonic/gin"
)

func Borrow(c *gin.Context) {
	_, user := GetId(GetToken(c))
	id := c.Param("id")

	if model.DB.Where("mochine_id = ?", id).
		Where("is_return = false").First(&model.McBorrow{}).RecordNotFound() {
		model.DB.Create(&model.McBorrow{
			PersonId:   user,
			MochineId:  id,
			BorrowTime: time.Now(),
			IsReturn:   false,
		})
		c.JSON(200, gin.H{
			"err": 0,
		})
	} else {
		c.JSON(200, gin.H{
			"err": 1,
			"msg": "Not Returned.",
		})
	}
}

func Return(c *gin.Context) {
	_, user := GetId(GetToken(c))
	id := c.Param("id")

	b := model.McBorrow{}
	if !model.DB.Where("person_id = ? and mochine_id = ? and is_return = false", user, id).First(&b).RecordNotFound() {
		b.IsReturn = true
		model.DB.Update(&b)
		c.JSON(200, gin.H{
			"err": 0,
		})
	} else {
		c.JSON(200, gin.H{
			"err": 1,
			"msg": "无此记录.",
		})
	}
}

func BorrowSearch(c *gin.Context) {
	sq := struct {
		Mochine  string
		Person   string
		Returned string
		Before   int64
		After    int64
	}{}
	var ps []model.McBorrow
	if err := c.ShouldBindQuery(&sq); err == nil {
		q := model.DB
		if sq.Mochine != "" {
			q = q.Where("mochine_id = ?", sq.Mochine)
		}
		if sq.Person != "" {
			q = q.Where("person_id = ?", sq.Person)
		}
		if sq.Returned == "true" {
			q = q.Where("is_return = ?", true)
		} else if sq.Returned == "false" {
			q = q.Where("is_return = ?", false)
		}
		if sq.Before > 0 {
			q = q.Where("borrow_time < ?", time.Unix(sq.Before, 0))
		}
		if sq.After > 0 {
			q = q.Where("borrow_time > ?", time.Unix(sq.After, 0))
		}

		q.Find(&ps)
		c.JSON(200, gin.H{
			"err":  0,
			"list": ps,
		})
	} else {
		c.JSON(404, gin.H{
			"err": 1,
			"msg": "查询错误",
		})
	}
}
