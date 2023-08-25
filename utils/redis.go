package utils

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdClient *redis.Client //指针
var duration = 30 * 24 * 60 * 60 * time.Second

type RedisClient struct {
}

func InitRedis() (*RedisClient, error) {
	rdClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &RedisClient{}, err
}

// *RedisClient包装三个方法
func (rdc *RedisClient) Set(key string, value any) error {
	return rdClient.Set(key, value, viper.GetDuration("redis.redisExpire")*time.Minute).Err()
}
func (rdc *RedisClient) Get(key string) (any, error) {
	return rdClient.Get(key).Result()
}
func (rdc *RedisClient) Delete(key ...string) error {
	return rdClient.Del(key...).Err()
}

func (rdc *RedisClient) GetExpireTime(key string) (time.Duration, error) {
	return rdClient.TTL(key).Result()
}
