package utils

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"time"
)

// Redis共享存储工具包
type RedisStore struct {
	Client     *redis.Client
	PreKey     string
	Expiration time.Duration
	Context    context.Context
}

func NewRedisStore(client *redis.Client, preKey string, expire time.Duration, context context.Context) *RedisStore {
	return &RedisStore{
		Client:     client,
		PreKey:     preKey,
		Expiration: expire,
		Context:    context,
	}
}

func (rs *RedisStore) Set(key string, val interface{}) {
	bytes, _ := json.Marshal(val)
	rs.Client.Set(rs.Context, rs.PreKey+key, bytes, rs.Expiration)
}

func (rs *RedisStore) Get(key string, obj interface{}) bool {
	val, err := rs.Client.Get(rs.Context, key).Result()
	if err != nil {
		return false
	} else {
		json.Unmarshal([]byte(val), obj)
		return true
	}
}
