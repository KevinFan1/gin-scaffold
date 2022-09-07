package internal

import (
	"bytes"
	"code/gin-scaffold/internal/global"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Insure(url string) {
	contentType := "application/json"
	data := make(map[string]string)
	data["username"] = "123"
	bytesData, _ := json.Marshal(data)

	ret, err := http.Post(url, contentType, bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		return
	}
	defer ret.Body.Close()

	content, err := ioutil.ReadAll(ret.Body)
	global.Logger.Info("post请求：", content)
}
