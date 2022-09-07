package middleware

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	r.Use(
		CORS(),
		Error(),
		IpLimiter(),
		//Auth(),
		Throttle(),
		ActionLog(),
		//JWTMiddleware(),
	)

}
