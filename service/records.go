package service

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/pagination"
	"code/gin-scaffold/models"

	"github.com/gin-gonic/gin"
)

type RecordsSrv struct {
}

var RecordService = new(RecordsSrv)

func (srv *RecordsSrv) List(c *gin.Context) (*pagination.Paginator, error) {
	var records []models.LoginRecord
	db := global.DB.Preload("User")
	paginator, err := pagination.Scan(c, db, records)
	return paginator, err
}
