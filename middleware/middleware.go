package middleware

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// cors 必须放在最前面, 记录时间
	// action log放后面，计算总共耗时
	r.Use(
		CORS(),
		Error(),
		IpLimiter(),
		Throttle(),
		ActionLog(),
	)

}
