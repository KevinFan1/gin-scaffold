package v1

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/pagination"
	"code/gin-scaffold/internal/utils"
	"code/gin-scaffold/internal/vo"
	"code/gin-scaffold/models"
	"code/gin-scaffold/schemas"
	"github.com/gin-gonic/gin"
)

// UserList 获取用户列表
func UserList(c *gin.Context) {
	//用户列表
	var users []models.User
	db := global.DB.Omit("Password").Preload("Role")
	paginator, err := pagination.Scan(c, db, users)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}
	vo.OkWithMsg(c, "查询用户列表成功", paginator)
}

// UserDetail 根据ID获取用户详情
func UserDetail(c *gin.Context) {
	userId := c.Query("user_id")
	user, _ := models.GetUserById(userId)
	vo.Ok(c, user)
}

// UserInfo 获取当前用户info
func UserInfo(c *gin.Context) {
	user := utils.GetCurrentUser(c)
	vo.Ok(c, user)
}

func UserAddition(c *gin.Context) {
	var userDto schemas.UserAddDto

	err := c.ShouldBindJSON(&userDto)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}

	err = global.DB.Create(&models.User{
		Username: userDto.Username,
		Password: userDto.Password,
		RoleId:   userDto.RoleId,
	}).Error
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}
	vo.OkWithMsg(c, "创建用户成功", nil)

}
