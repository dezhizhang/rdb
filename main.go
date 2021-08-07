package main

import (
	"context"
	"log"
	"rdb/driver"
)

func main() {
	ctx := context.Background()
	ret := driver.Redis().Get(ctx,"name")
	val,err := ret.Result()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(val)
}