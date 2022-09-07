package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5ByStr(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5ByBytes(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
