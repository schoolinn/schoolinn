package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Edu struct
{
	School string
	Professional string
	Level string
	Dec string
	Time string
}

type User struct
{
	Name string
	Edu_list []Edu
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")
	var users []User 
	c.Find(&bson.M{"_id": bson.ObjectIdHex("5a657f7c0dd364221c83784b")}).All(&users) 
	fmt.Println(users)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}