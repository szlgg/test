package utils

import (
	"crypto/md5"
	"fmt"
	"math/rand"
)

// MD5加密
func MustEncrypt(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

// 生成唯一的字符串
func RandomString() string {
	n := 16
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return s
}
