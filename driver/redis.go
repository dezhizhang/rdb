package driver

import (
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)


var redisClient * redis.Client
var redisClientOnce sync.Once

func Redis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB: 0,
		PoolSize: 15,
		MinIdleConns: 10,
		DialTimeout: 5 * time.Second,
		ReadTimeout: 3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})
	return redisClient
}

func init() {
	Redis()
}
