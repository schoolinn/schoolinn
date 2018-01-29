package que

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"schoolinn/mgo"
)

type Que struct
{
	Answer_num,Author,Time,Title,Dec string
	Answer    interface{}      `bson:"answer"`
	Tags    interface{}      `bson:"tags"`
}

func Cb( g *gin.Context ){

	g.Header("Access-Control-Allow-Origin", "*")

	c := mgo.GetDataBase().C("question")
	var que []Que 
	c.Find(nil).Skip(0).Limit(15).All(&que) 
	fmt.Println(que)
	s := mgo.GetMgo()
	s.Close()
	
	g.JSON(200, gin.H{
		"data": que,
	})
}