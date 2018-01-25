package tags

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"schoolinn/mgo"
)

type Tag struct
{
	Timg,Id,Name string
}

func Cb( g *gin.Context ){

	g.Header("Access-Control-Allow-Origin", "*")

	c := mgo.GetDataBase().C("tags")
	var tag []Tag 
	c.Find(nil).All(&tag) 
	fmt.Println(tag)
	s := mgo.GetMgo()
	s.Close()
	
	g.JSON(200, gin.H{
		"data": tag,
	})
}