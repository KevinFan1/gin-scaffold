package middleware

import (
	"code/gin-scaffold/internal/acs"
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/settings"
	"code/gin-scaffold/internal/utils"
	"code/gin-scaffold/internal/vo"
	"code/gin-scaffold/models"
	"code/gin-scaffold/schemas"
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var (
	identityKey = "id"
	sessionKey  = "session_id"

	ErrAuthentication = errors.New("账号或密码错误")
	ErrMissingParams  = errors.New("缺少用户名或密码")
)

type UserJwtClaims struct {
	ID        any
	SessionId any
}

func JWTMiddleWareGenerator() *jwt.GinJWTMiddleware {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test jwt",
		Key:         []byte(settings.Setting.SystemBaseConfig.SecretKey),
		Timeout:     time.Hour * 2,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			// 登录的时候设置payload
			if v, ok := data.(*UserJwtClaims); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					sessionKey:  v.SessionId,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			// 每次进来校验该参数,负责给 Authorizator Handler 传递内容
			claims := jwt.ExtractClaims(c)
			userId := int(claims[identityKey].(float64))

			user, err := models.GetUserById(userId)

			if err != nil {
				c.Abort()
				vo.Unauthorized(c)
				return nil
			}

			// 设定当前request user
			c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "user", user))

			// TODO: 查询数据库操作,查询角色
			return &UserJwtClaims{
				ID:        claims[identityKey],
				SessionId: claims[sessionKey],
			}

		},

		Authenticator: func(c *gin.Context) (interface{}, error) {
			// 登录 handler 校验
			var loginDto schemas.LoginDto
			err := c.ShouldBindJSON(&loginDto)

			if err != nil {
				return nil, ErrMissingParams
			}

			var user models.User

			if err = global.DB.Where(&models.User{Username: loginDto.Username}).Select("id", "password").First(&user).Error; err != nil {
				return nil, ErrAuthentication
			}

			record := models.LoginRecord{
				UserId:  user.ID,
				Agent:   c.Request.UserAgent(),
				Ip:      c.RemoteIP(),
				LoginAt: models.JSONTime{Time: time.Now()},
			}

			// 校验密码
			if user.Password != loginDto.Password {
				record.Status = strconv.Itoa(http.StatusUnauthorized)
				record.Code = "-1"
				record.Response = ErrAuthentication.Error()

				global.DB.Create(&record)

				return nil, ErrAuthentication
			}
			record.Status = strconv.Itoa(http.StatusOK)
			record.Code = "0"
			record.Response = "登录成功"
			global.DB.Create(&record)

			// TODO: save token(uuid) to redis
			return &UserJwtClaims{
				ID:        user.ID,
				SessionId: utils.NewUUID(),
			}, nil
		},

		Authorizator: func(data interface{}, c *gin.Context) bool {
			// 权限管理, 结合casbin
			v, ok := data.(*UserJwtClaims)
			if !ok || v == nil {
				return false
			}
			path := c.Request.URL.Path
			method := c.Request.Method
			user := utils.GetCurrentUser(c)

			if user.Role == nil {
				return false
			}

			// 去掉前缀, 更改后不需要修改数据库
			path = strings.Replace(path, settings.Setting.SystemBaseConfig.ApiPrefix, "", 1)
			return acs.CheckPermission(user.Role.Code, path, method)

		},

		Unauthorized: func(c *gin.Context, code int, message string) {
			// 上述函数校验失败返回内容
			switch code {
			case http.StatusForbidden:
				vo.Forbidden(c)
			case http.StatusUnauthorized:
				vo.Unauthorized(c)
			default:
				vo.Unauthorized(c)
			}
		},
		// 可以header query cookie三个地方获取凭证信息
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// token 前缀必须是Bearer，且需要用空格隔开token
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		panic("Jwt init error: " + err.Error())
	}
	return middleware
}

func JWTMiddleware() gin.HandlerFunc {
	return JWTMiddleWareGenerator().MiddlewareFunc()
}
