package driver

import (
	"context"
	"fmt"
)

type Locker struct {
	key string
}

func NewLocker(key string) *Locker {
	return &Locker{key: key}
}

func (this *Locker) Locker() *Locker {
	boolcmd := redisClient.SetNX(context.Background(),this.key,1,0)
	if ok,err := boolcmd.Result();err != nil || !ok{
		panic(fmt.Sprintf("lock error width key:%s",this.key))
	}
	return this
}

func (this *Locker) Unlock() {
	redisClient.Del(context.Background(),this.key)
}
