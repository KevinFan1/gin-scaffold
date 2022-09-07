package v1

import (
	"code/gin-scaffold/internal/utils"
	"code/gin-scaffold/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type name struct {
	Hello string
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World! :)")
	//job := asyncworker.Job{Payload: name{
	//	Hello: "123123",
	//}}
	//// 发送保险订单任务 异步
	//fmt.Println("job payload:", job)

	//context.Abort()

	//asyncworker.JobQueue <- asyncworker.Job{}
	//context.String(http.StatusOK, "Hello World")
	//fmt.Printf("异步任务结果：%v\n", asyncResult)
}

func XmlReader(c *gin.Context) {
	var xmlData schemas.XMLData

	err := c.BindXML(&xmlData)

	if err != nil {
		c.String(http.StatusBadRequest, "fail to parse post data")
		return
	}

	if err = utils.CheckCode(xmlData); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("%v", err)+"    IP:"+c.RemoteIP())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "this is ok message",
	})
}
