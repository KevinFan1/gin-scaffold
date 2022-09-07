package controller

import v1 "code/gin-scaffold/controller/v1"

var (
	UserListController     = v1.UserList
	UserDetailController   = v1.UserDetail
	UserInfoController     = v1.UserInfo
	UserAdditionController = v1.UserAddition

	CreateCasbinRuleController = v1.CreateCasbinRule
	DeleteCasbinRuleController = v1.DeleteCasbinRule
)
