package middleware

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/vo"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	lgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	lredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func Throttle() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 10 * time.Second,
		Limit:  10,
	}

	store, err := lredis.NewStoreWithOptions(global.RedisClient, limiter.StoreOptions{
		Prefix:   "limiter_gin",
		MaxRetry: 3,
	})
	if err != nil {
		global.Logger.Fatalf("Throttle catch err: %v", err)
	}

	selfLimiter := limiter.New(store, rate)

	middleware := &lgin.Middleware{
		Limiter: selfLimiter,
		OnError: lgin.DefaultErrorHandler,
		OnLimitReached: func(c *gin.Context) {
			vo.ThrottleLimit(c)
		},
		KeyGetter:   lgin.DefaultKeyGetter,
		ExcludedKey: nil,
	}

	return middleware.Handle
}
