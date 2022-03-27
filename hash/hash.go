package hash

import (
	"crypto/md5"
	"fmt"
	"github.com/satori/go.uuid"
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

func GenUUID() string {
	var err error
	u1 := uuid.Must(uuid.NewV4(), err)
	return u1.String()
}
