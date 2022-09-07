package v1

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/pagination"
	"code/gin-scaffold/internal/vo"
	"code/gin-scaffold/models"
	"github.com/gin-gonic/gin"
)

func LoginRecordList(c *gin.Context) {
	var records []models.LoginRecord
	db := global.DB.Preload("User")
	paginator, err := pagination.Scan(c, db, records)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}
	vo.OkWithMsg(c, "查询登录日志成功", paginator)
}
