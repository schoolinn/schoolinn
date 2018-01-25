package tag_art

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"schoolinn/mgo"
	"gopkg.in/mgo.v2/bson"
)

type Art struct
{
	Title, Content string
}

func Cb( g *gin.Context ){

	g.Header("Access-Control-Allow-Origin", "*")

	c := mgo.GetDataBase().C("article")
	var art []Art
	fmt.Println(g.Param("id"))
	c.Find(&bson.M{"tag.id": g.Param("id")}).Skip(0).Limit(15).All(&art) 
	fmt.Println(art)
	s := mgo.GetMgo()
	s.Close()
	
	g.JSON(200, gin.H{
		"data": art,
	})
}