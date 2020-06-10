package model

import (
	"time"
)

type McBorrow struct {
	Id         int
	PersonId   string `gorm:"not null;type:varchar(20)"`
	MochineId  string `gorm:"not null;type:varchar(20)"`
	BorrowTime time.Time
	IsReturn   bool `gorm:"default:false"`
}
