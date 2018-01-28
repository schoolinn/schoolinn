package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"schoolinn/mgo"
)

type Art struct
{
	Title,Content,Dec,Time,Utime string
	Tag    interface{}      `bson:"tag"`
}

func Cb( g *gin.Context ){

	g.Header("Access-Control-Allow-Origin", "*")

	c := mgo.GetDataBase().C("article")
	var art []Art 
	c.Find(nil).Skip(0).Limit(15).All(&art) 
	fmt.Println(art)
	s := mgo.GetMgo()
	s.Close()
	
	g.JSON(200, gin.H{
		"data": art,
	})
}