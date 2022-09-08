package service

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/pagination"
	"code/gin-scaffold/models"
	"code/gin-scaffold/schemas"

	"github.com/gin-gonic/gin"
)

type UserSrv struct {
}

var UserService = new(UserSrv)

// List 获取用户列表
func (srv *UserSrv) List(c *gin.Context) (*pagination.Paginator, error) {
	//用户列表
	var users []models.User
	db := global.DB.Omit("Password").Preload("Role")
	paginator, err := pagination.Scan(c, db, users)
	if err != nil {
		return nil, err
	}
	return paginator, nil
}

// Detail 根据user_id获取用户信息
func (srv *UserSrv) Detail(c *gin.Context) *models.User {
	userId := c.Query("user_id")
	user, _ := models.GetUserById(userId)
	return user
}

// Create 创建用户
func (srv UserSrv) Create(c *gin.Context) error {
	var userDto schemas.UserAddDto
	// 绑定用户数据
	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		return err
	}

	err = global.DB.Create(&models.User{
		Username: userDto.Username,
		Password: userDto.Password,
		RoleId:   userDto.RoleId,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
