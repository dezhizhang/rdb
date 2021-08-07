package main

import (
	"github.com/gin-gonic/gin"
	"rdb/model"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		str := c.Query("id")
		id,_:= strconv.Atoi(str)
		user := model.NewUser().LoadById(id)
		c.JSON(200,gin.H{"code":200,"result":user})
	})
	r.Run()
}
