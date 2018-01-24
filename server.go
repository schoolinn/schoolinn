package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"schoolinn/mgo"
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
	
	c := mgo.GetDataBase().C("people")
	var users []User 
	c.Find(nil).All(&users) 
	fmt.Println(users)
	s := mgo.GetMgo()
	s.Close()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": users,
		})
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}