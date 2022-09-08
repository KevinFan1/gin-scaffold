package models

import (
	"code/gin-scaffold/internal/global"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// User todo: password encrypt
type User struct {
	BaseModel
	Username  string   `json:"username" gorm:"column:username;type:varchar(32);comment:用户名;not null;default:"`
	Password  string   `json:"-" gorm:"column:password;type:varchar(255);comment:密码;not null;default:;"`
	RoleId    uint     `json:"-" gorm:"column:role_id;comment:角色id;"`
	Role      *Role    `json:"role,omitempty" gorm:"foreignKey:RoleId;"`
	LastLogin JSONTime `json:"last_login" gorm:"column:last_login;comment:最后登录时间"`
	IsDisable bool     `json:"is_disable" gorm:"column:is_disable;type:tinyint;comment:是否禁用;default:0;"`
	IsDel     bool     `json:"is_del" gorm:"column:is_del;type:tinyint;comment:是否删除;default:0;"`
}

func (u User) TableName() string {
	return "t_user"
}

func GetUserById[T int | string](userId T) (user *User, err error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("user:%v", userId)

	val, err := global.RedisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		err = global.DB.Omit("Password").Where("id=?", userId).Preload("Role").First(&user).Error
		marshal, _ := json.Marshal(user)
		global.RedisClient.Set(ctx, cacheKey, marshal, 10*time.Minute)
	} else if err != nil {
		return nil, err
	} else {
		_ = json.Unmarshal([]byte(val), &user)
	}
	return user, err
}
