package redis

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

func InitRedSync(client redis.UniversalClient) *redsync.Redsync {
	// 创建 Redsync 的 Redis 连接池
	pool := goredis.NewPool(client)

	// 创建 Redsync 实例
	return redsync.New(pool)
}
