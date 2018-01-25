package user

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

type Worklist struct
{
	Job string
	Cname string
	Dec string
	Time string
}

type Skill struct
{
	Name string
    Level string
}

type Works struct
{
	Imgs []string
	Dec string
	Name string
}

type User struct
{
	Name,Dec,Work_price,Work_time,Address,Avatar string
	Edu_list []Edu
	Skill_list []Skill
	Works []Works
	Work_list []Worklist
}

func Cb( g *gin.Context ){

	g.Header("Access-Control-Allow-Origin", "*")

	c := mgo.GetDataBase().C("people")
	var users []User 
	c.Find(nil).Skip(0).Limit(15).All(&users) 
	fmt.Println(users)
	s := mgo.GetMgo()
	s.Close()
	
	g.JSON(200, gin.H{
		"data": users,
	})
}
