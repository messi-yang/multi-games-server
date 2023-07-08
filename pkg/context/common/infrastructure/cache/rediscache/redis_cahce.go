package rediscache

import (
	"context"
	"time"

	"github.com/dum-dum-genius/zossi-server/pkg/context/common/infrastructure/redisclient"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type RedisCacher interface {
	Set(key string, value string, duration time.Duration) error
	Get(key string) (value string, err error)
	Del(key string) (err error)
	Scan(pattern string) (values []string, err error)
}

type redisCache struct {
	redisClient *redis.Client
	redisCache  *cache.Cache
}

func NewRedisCacher() RedisCacher {
	return &redisCache{
		redisClient: redisclient.GetRedisClient(),
		redisCache: cache.New(&cache.Options{
			Redis: redisclient.GetRedisClient(),
		}),
	}
}

func (redisCache *redisCache) Set(key string, value string, duration time.Duration) error {
	return redisCache.redisCache.Set(&cache.Item{
		Ctx:   context.TODO(),
		Key:   key,
		Value: value,
		TTL:   duration,
	})
}

func (redisCache *redisCache) Get(key string) (value string, err error) {
	if err = redisCache.redisCache.Get(context.TODO(), key, &value); err != nil {
		return value, err
	}
	return value, nil
}

func (redisCache *redisCache) Del(key string) (err error) {
	return redisCache.redisCache.Delete(context.TODO(), key)
}

func (redisCache *redisCache) Scan(pattern string) (values []string, err error) {
	values = make([]string, 0)
	iter := redisCache.redisClient.Scan(context.TODO(), 0, pattern, 200).Iterator()
	for iter.Next(context.TODO()) {
		playerKey := iter.Val()
		playerVal, err := redisCache.Get(playerKey)
		if err != nil {
			return values, err
		}
		values = append(values, playerVal)
	}
	return values, nil
}
