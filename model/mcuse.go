package model

import (
	"time"
)

type McUse struct {
	Id        int    `gorm:""`
	UserId    string `gorm:"not null;type:varchar(20)"`
	MochineId string `gorm:"not null;type:varchar(20)"`
	UseTime   time.Time
}
