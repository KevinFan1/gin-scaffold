package utils

import (
	"code/gin-scaffold/schemas"
	"errors"
	"fmt"
)

// genSign 生成sign签名
func genSign(xmlData schemas.XMLData) string {
	encryptStr := xmlData.Head.UUID + xmlData.Head.CRequestType + xmlData.Head.CBusiChnl
	return Md5ByStr(encryptStr)
}

// CheckCode 检测xml data签名是否正确
func CheckCode(xmlData schemas.XMLData) (err error) {
	rawCode := xmlData.Head.CheckCode
	signCode := genSign(xmlData)
	if rawCode != signCode {
		return errors.New(fmt.Sprintf("签名错误,原code:%v,计算code:%v", rawCode, signCode))
	}
	return nil
}
