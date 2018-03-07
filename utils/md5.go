package utils

import (
	"crypto/md5"
	"fmt"
)

//md5
func MD5Sum(data string) string {
	if len(data) == 0 {
		return ""
	}
	ret := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", ret)
}