package model

type Mochine struct {
	Id       string `gorm:"PRIMARY_KEY;not null;type:varchar(20)";json:"id"`
	AdminId  string `gorm:"not null;varchar(20)";json:"adminId"`
	Kind     string `gorm:"not null";json:"kind"`
	LifeYear int    `gorm:"not null";json:"lifeYear"`
}
