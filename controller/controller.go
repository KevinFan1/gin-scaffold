package controller

import (
	v1 "code/gin-scaffold/controller/v1"
)

var (
	GetUserList   = v1.UserList
	GetUserDetail = v1.UserDetail
	GetUserInfo   = v1.UserInfo
	CreateUser    = v1.UserAddition

	CreateCasbinRuleController = v1.CreateCasbinRule
	DeleteCasbinRuleController = v1.DeleteCasbinRule

	LoginRecordListController = v1.LoginRecordList
)
