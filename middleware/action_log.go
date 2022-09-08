package middleware

import (
	"bytes"
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/utils"
	"code/gin-scaffold/internal/vo"
	"code/gin-scaffold/models"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func ActionLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不记录GET操作
		if c.Request.Method == "GET" {
			fmt.Println("不记录get操作")
			c.Next()
			return
		}

		blw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		// 如果有标记，则记录日志

		// 记录request body
		bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		c.Next()

		logName := c.Request.Context().Value("log")
		user := utils.GetCurrentUser(c)
		if logName == nil || user == nil {
			fmt.Println("不需要记录")
			return
		}

		var cr vo.CommonResponse
		err := json.Unmarshal(blw.body.Bytes(), &cr)
		if err != nil {
			global.Logger.Errorf("解析common reponse错误%v\n", err)
			return
		}

		method := c.Request.Method
		path := c.Request.URL.Path

		content := ""
		response := ""

		if method == "GET" {
			content = c.Request.URL.RawQuery
			response = cr.Msg
		} else {
			content = string(bodyBytes)
			response = blw.body.String()
		}
		//计算总耗时
		cost := time.Now().UnixMilli() - c.Request.Context().Value("start_time").(int64)
		record := models.OperationRecord{
			UserId:   utils.GetCurrentUser(c).ID,
			Agent:    c.Request.UserAgent(),
			Ip:       c.ClientIP(),
			Path:     path,
			Method:   method,
			Content:  content,
			Status:   fmt.Sprintf("%v", blw.Status()),
			Code:     fmt.Sprintf("%v", cr.Code),
			Response: response,
			Cost:     fmt.Sprintf("%v", cost),
		}
		// todo:任务池消费任务
		go global.DB.Create(&record)
	}
}
