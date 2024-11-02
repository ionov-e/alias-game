package database

import (
	"github.com/redis/go-redis/v9"
)

type RedisClient = redis.Client

func NewClient(options *redis.Options) *RedisClient {
	return redis.NewClient(options)
}

func NewClientLocal() *RedisClient {
	return NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
