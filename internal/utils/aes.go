package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
)

type AesCrypt struct {
	Key []byte
}

////////////////////////  CBC  ////////////////////////

// pkcs5Padding 添加码
func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)

}

// pkcs5UnPadding 去除密文的码
func pkcs5UnPadding(rawDataByte []byte) []byte {
	dataLength := len(rawDataByte)
	unpadding := int(rawDataByte[dataLength-1])
	return rawDataByte[:dataLength-unpadding]
}

func (e AesCrypt) CBCEncrypt2Byte(RawData []byte) (encrypted []byte) {
	block, _ := aes.NewCipher(e.Key[:16])
	// 获取密钥的长度
	blockSize := block.BlockSize()
	// 补全码
	RawData = pkcs5Padding(RawData, blockSize)
	iv := e.Key[:blockSize]
	// CBC加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 加密的结果 byte
	encrypted = make([]byte, len(RawData))
	// 加密数据，
	blockMode.CryptBlocks(encrypted, RawData)
	return encrypted
}

func (e AesCrypt) CBCEncrypt2Str(RawData []byte) (encryptedStr string) {
	// 加密成byte后，转化成string类型返回
	encrypted := e.CBCEncrypt2Byte(RawData)
	return hex.EncodeToString(encrypted)

}

func (e AesCrypt) CBCDecryptByte(encrypted []byte) (string, error) {
	block, _ := aes.NewCipher(e.Key[:16])
	// 获取密钥的长度
	blockSize := block.BlockSize()
	if len(encrypted)%blockSize != 0 {
		return "", errors.New("解密错误,请联系管理员")
	}

	// CBC解密模式
	blockMode := cipher.NewCBCDecrypter(block, e.Key[:blockSize])
	// 创建数组
	decrypted := make([]byte, len(encrypted))
	// 解密
	blockMode.CryptBlocks(decrypted, encrypted)
	// 去除补全码
	decrypted = pkcs5UnPadding(decrypted)

	return string(decrypted), nil
}

func (e AesCrypt) CBCDecryptStr(encrypted string) (string, error) {
	encryptedByte, _ := hex.DecodeString(encrypted)
	decryptStr, err := e.CBCDecryptByte(encryptedByte)
	if err != nil {
		return "", err
	}

	return decryptStr, nil
}

////////////////////////  CBC  ////////////////////////
