package middleware

import (
	"bytes"
	"code/gin-scaffold/internal/vo"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
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
		blw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		//先放行,如果有标记，则记录日志
		c.Next()

		method := c.Request.Method
		path := c.Request.URL.Path
		logName := c.Request.Context().Value("log")

		if method != "GET" && logName != nil {
			fmt.Println(path)
			//fmt.Println(ioutil.ReadAll(c.Request.Body))
		}

		//fmt.Println(method)
		//fmt.Println(path)

		var common vo.CommonResponse
		_ = json.Unmarshal([]byte(blw.body.String()), &common)
		//fmt.Println(common)
		//fmt.Println(c.Request.Context().Value("log"))
		//fmt.Println(utils.GetCurrentUser(c))
	}
}
