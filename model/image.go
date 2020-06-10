package model

type Image struct {
	Person Person `gorm:"ForeignKey:Id;Association_ForeignKey:Id"`
	Id     string `gorm:"PRIMARY_KEY;not null;type:varchar(20)"`
	File   []byte
	Type   string
}

var TypeMap = map[string]string{
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"gif":  "image/gif",
	"bmp":  "image/bmp",
	"webp": "image/webp",
}
