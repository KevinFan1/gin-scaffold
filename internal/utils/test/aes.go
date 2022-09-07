package main

import (
	"code/gin-scaffold/internal/utils"
	"encoding/hex"
	"fmt"
)

func main() {
	crypt := utils.AesCrypt{Key: []byte("l)wtsz8j92mz$d4mhio(1o_!64ivagf$_c#n5r&cuh^=g_&1(=")}
	data := crypt.CBCEncrypt2Byte([]byte("pj:login:token:1"))
	fmt.Println(crypt.CBCDecryptByte(data))

	ss := crypt.CBCEncrypt2Str([]byte("pj:login:token:1"))
	decodeString, err := hex.DecodeString(ss)
	if err != nil {
		return
	}
	fmt.Println(crypt.CBCDecryptByte(decodeString))
}
