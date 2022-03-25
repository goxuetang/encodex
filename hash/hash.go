package hash

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string, salt string) string {
	if str == "" {
		return ""
	}
	h := md5.New()
	h.Write([]byte(str+salt))
	re := h.Sum(nil)
	md5str := fmt.Sprintf("%x\n", re)
	return md5str
}