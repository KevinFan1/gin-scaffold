package vo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommonResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

const (
	SUCCESS = 0
	FAIL    = -1
)

func result(c *gin.Context, status int, code int, data any, msg string) {
	c.JSON(status, CommonResponse{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Ok(c *gin.Context, data any) {
	result(c, http.StatusOK, SUCCESS, data, "操作成功")
}

func OkWithMsg(c *gin.Context, msg string, data ...any) {
	result(c, http.StatusOK, SUCCESS, data, msg)
}

func Unauthorized(c *gin.Context) {
	result(c, http.StatusUnauthorized, FAIL, nil, "凭证验证失败,请重新登录")
}

func Forbidden(c *gin.Context) {
	result(c, http.StatusForbidden, FAIL, nil, "无权访问,请联系管理员进行授权")
}

func FailWithMsg(c *gin.Context, msg string) {
	result(c, http.StatusBadRequest, FAIL, nil, msg)
}

func ServerError(c *gin.Context) {
	result(c, http.StatusInternalServerError, FAIL, nil, "Internal Server Error|服务器内部错误")
}

func ThrottleLimit(c *gin.Context) {
	result(c, http.StatusTooManyRequests, FAIL, nil, "访问频繁,请稍后重试")
}
