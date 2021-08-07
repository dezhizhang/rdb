package main

import (
	"github.com/gin-gonic/gin"
	"rdb/driver"
)

func main() {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		defer func() {
			if e := recover();e != nil {
				c.AbortWithStatusJSON(400,gin.H{"message":e})
			}
		}()
		c.Next()
	})
	r.GET("/", func(c *gin.Context) {
		locker := driver.NewLocker("locker").Locker()
		defer locker.Unlock()
		c.JSON(200,gin.H{"message":"ok"})
	})

	r.Run()
}
