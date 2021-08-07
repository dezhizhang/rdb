package driver
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)


var redisClient * redis.Client
var redisClientOnce sync.Once

func Redis() *redis.Client {
	redisClientOnce.Do(func() {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
			PoolSize:15,
		})
		pong,err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatal(fmt.Errorf("connect error:%s",err))
		}
		log.Fatal(pong)
	})

	return redisClient

}
