package driver
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)


var redisClient * redis.Client
var redisClientOnce sync.Once

func Redis() *redis.Client {
	redisClientOnce.Do(func() {
		redisClient := redis.NewClient(&redis.Options{
			Addr: "127.0.0.1:6379",
			Password: "",
			DB: 0,
			PoolSize: 15,
			MinIdleConns: 10,
			DialTimeout: 5 * time.Second,
			ReadTimeout: 3 * time.Second,
			WriteTimeout: 3 * time.Second,
		})
		pong,err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal(fmt.Errorf("连接失败:%err",err))
			return
		}
		log.Fatal(pong)
	})
	return redisClient

}
