package util

import (
	"crypto/md5"
	"fmt"
)

func Password2MD5(p string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(p)))
}
