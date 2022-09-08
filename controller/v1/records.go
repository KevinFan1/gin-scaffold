package v1

import (
	"code/gin-scaffold/internal/vo"

	"github.com/gin-gonic/gin"
)

func LoginRecordList(c *gin.Context) {
	paginator, err := recordService.List(c)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}
	vo.OkWithMsg(c, "查询登录日志成功", paginator)
}
