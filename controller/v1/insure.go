package v1

import (
	"code/gin-scaffold/internal/vo"
	"code/gin-scaffold/internal/worker"
	"code/gin-scaffold/schemas"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func Insure(c *gin.Context) {
	// 解析数据
	var packetData schemas.Packet

	err := c.ShouldBindXML(&packetData)
	if err != nil {
		resp := vo.CommonResponse{Code: -1, Msg: fmt.Sprintf("解析xml错误.原因:%v", err)}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	//err = gredis.SSet("123", "!23", 10*time.Minute)
	//
	//if err != nil {
	//	logging.SugarLogger.Errorf("保存redis失败, err: %v", err)
	//} else {
	//	logging.SugarLogger.Info("保存redis成功")
	//}

	//发送任务
	work := worker.Job{Payload: name{
		Hello: "123123",
	}}
	//start := time.Now()
	// 发送保险订单任务 异步
	worker.JobQueue <- work
	//fmt.Printf("耗费时间:%v \n", time.Now().Sub(start))
	c.JSON(http.StatusOK, vo.CommonResponse{
		Code: 0, Msg: "解析数据成功", Data: schemas.InsureResponse{No: uuid.New().String(), Data: packetData},
	})

}
