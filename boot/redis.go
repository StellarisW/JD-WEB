package boot

import (
	"context"
	"github.com/go-redis/redis/v8"
	"main/app/global"
	"time"
)

var rdb *redis.Client

// RedisSetup Initialize the Redis instance
func RedisSetup() {
	config := g.Config.Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
		PoolSize: config.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		g.Logger.Fatalf("Connect to Redis instance failed, err: %v\n", err)
		//log.set
	}
	g.Redis = rdb

	g.Logger.Info("Initialize Redis instance successfully!")
}
