package gredis

import (
	"code/gin-scaffold/internal/global"
	"code/gin-scaffold/internal/settings"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type RedisHelper struct {
	*redis.Client
}

var (
	ctx = context.Background()
)

func GetRedisClient() *redis.Client {
	return global.RedisClient
}

// NewRedisClient 初始化一个redis client
func NewRedisClient() {
	global.RedisClient = redis.NewClient(&redis.Options{
		// 连接信息配置
		Addr:     settings.Setting.RedisConfig.Host + ":" + settings.Setting.RedisConfig.Port,
		Password: settings.Setting.RedisConfig.Password,
		DB:       settings.Setting.RedisConfig.DB,
		//重试策略配置
		MaxRetries:      3,
		MinRetryBackoff: 8 * time.Millisecond,
		MaxRetryBackoff: 512 * time.Millisecond,
		//超时配置
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		// 连接池 配置
		PoolSize:           settings.Setting.RedisConfig.PoolSize,
		MinIdleConns:       8,
		MaxConnAge:         10 * time.Minute,
		PoolTimeout:        4 * time.Minute,
		IdleTimeout:        time.Duration(settings.Setting.RedisConfig.MaxIdle) * time.Minute,
		IdleCheckFrequency: 1 * time.Minute,

		// 钩子函数，在建立链接的时候调用
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			log.Println("建立Redis链接:", cn)
			return nil
		},
	})
}

// SetUp 普通单机连接
func SetUp() {
	//初始化redis client
	NewRedisClient()
	pong, err := global.RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis连接失败. err:", zap.Error(err))
	} else {
		log.Println("Redis链接成功, ping响应:", zap.String("pong", pong))
	}
}

//哨兵模式client
func initSentinelClient() (err error) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379"},
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

//集群redis
func initClusterClient() (err error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001"},
	})
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

// SSet 相关操作
func SSet(key string, val interface{}, expire time.Duration) error {
	start := time.Now()
	value, err := json.Marshal(val)
	if err != nil {
		fmt.Println("RedisClient Set value with error:", err)
		return err
	}
	err = global.RedisClient.Set(ctx, key, value, expire).Err()
	if err != nil {
		return err
	}
	fmt.Println("保存redis时间:", time.Now().Sub(start))
	return nil
}

func SGet(key string) (string, error) {
	return global.RedisClient.Get(ctx, key).Result()
}

/* HSet Hash类型相关操作*/

func HSet(key string, id string, val interface{}) {
	global.RedisClient.HSet(ctx, key, id, val)
}
