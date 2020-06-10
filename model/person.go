package model

const (
	DefaultLevel = "default"
	AdminLevel   = "admin"
)

type Person struct {
	Id       string `gorm:"PRIMARY_KEY;not null;type:varchar(20)";json:"id"`
	Name     string `gorm:"not null;type:varchar(20)";json:"name"`
	Level    string `gorm:"default:'default'";json:"level"`
	Password string `gorm:"not null";json:"password"`
}
