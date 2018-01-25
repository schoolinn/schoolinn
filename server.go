package main

import (
	"github.com/gin-gonic/gin"
	"schoolinn/user"
	"schoolinn/article"
	"schoolinn/tags"
	"schoolinn/tag_art"
)

func main() {
	
	r := gin.Default()
	r.GET("/ping", user.Cb)
	r.GET("/article", article.Cb)
	r.GET("/tag", tags.Cb)
	r.GET("/tag_art/:id", tag_art.Cb)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}