package models

import (
	"code/gin-scaffold/internal/global"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// User todo: password encrypt
type User struct {
	BaseModel
	Username string `gorm:"column:username;type:varchar(32);comment:用户名;not null;default:" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);comment:密码;not null;default:;" json:"password"`
	RoleId   uint   `gorm:"column:role_id;comment:角色id;" json:"role_id"`
	Role     *Role  `gorm:"foreignKey:RoleId;" json:"role,omitempty"`
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
