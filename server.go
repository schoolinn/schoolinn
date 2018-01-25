package main

import (
	"github.com/gin-gonic/gin"
	"schoolinn/user"
)

func main() {
	
	r := gin.Default()
	r.GET("/ping", user.Cb)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}