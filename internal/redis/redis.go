package redis

import (
	"context"
	"crmeb_go/internal/conf"
	"github.com/redis/go-redis/v9"

	"time"
)

func InitRedis() (*redis.Client, error) {
	// 创建 Redis 客户端
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Config.Redis.Addr,
		Password:     conf.Config.Redis.Password, // 如果没有密码，使用空字符串
		DB:           conf.Config.Redis.DB,
		Username:     conf.Config.Redis.Username,
		PoolSize:     16, // 连接池大小
		MinIdleConns: 3,  // 最小空闲连接数
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}
	return client, nil
}
