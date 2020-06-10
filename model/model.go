package model

import (
	"sx/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDB(url string) (*gorm.DB, error) {
	d, err := gorm.Open("sqlite3", url)
	if err != nil {
		return nil, err
	}
	DB = d
	return d, nil
}

func AddPerson(id string, name string, level string, password string) (ok bool) {
	if DB.First(Person{Id: id}).RecordNotFound() {
		DB.Create(Person{
			id,
			name,
			level,
			util.Password2MD5(password),
		})
		return true
	} else {
		return false
	}
}

func LoginCheck(id string, password string) (ok bool, level string) {
	p := Person{Id: id, Password: util.Password2MD5(password)}
	if DB.First(&p).RecordNotFound() {
		return false, ""
	} else {
		return true, p.Level
	}
}
