package captcha

import (
	"context"
	"go.uber.org/zap"
	g "main/app/global"
	"time"
)

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "captcha_",
		Context:    context.Background(),
	}
}

func (rs *RedisStore) Set(id string, value string) error {
	err := g.Redis.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		g.Logger.Error("RedisStoreSetError!", zap.Error(err))
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := g.Redis.Get(rs.Context, key).Result()
	if err != nil {
		g.Logger.Error("RedisStoreGetError!", zap.Error(err))
		return ""
	}
	if clear {
		err := g.Redis.Del(rs.Context, key).Err()
		if err != nil {
			g.Logger.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
