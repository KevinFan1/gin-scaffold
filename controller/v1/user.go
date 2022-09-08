package v1

import (
	"code/gin-scaffold/internal/utils"
	"code/gin-scaffold/internal/vo"
	"github.com/gin-gonic/gin"
)

// UserList 获取用户列表
func UserList(c *gin.Context) {
	//用户列表
	paginator, err := userService.List(c)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}
	vo.OkWithMsg(c, "查询用户列表成功", paginator)
}

// UserDetail 根据ID获取用户详情
func UserDetail(c *gin.Context) {
	user := userService.Detail(c)
	vo.Ok(c, user)
}

// UserInfo 获取当前用户info
func UserInfo(c *gin.Context) {
	user := utils.GetCurrentUser(c)
	vo.Ok(c, user)
}

// UserAddition 添加用户
func UserAddition(c *gin.Context) {
	err := userService.Create(c)
	if err != nil {
		vo.FailWithMsg(c, err.Error())
		return
	}
	vo.OkWithMsg(c, "创建用户成功", nil)

}
