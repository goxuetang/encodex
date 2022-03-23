package hash

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	re := h.Sum(nil)
	md5str := fmt.Sprintf("%x\n", re)

	return md5str
}