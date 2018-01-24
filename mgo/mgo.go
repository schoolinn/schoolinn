package mgo

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var database *mgo.Database



func init(){
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	
	database = session.DB("test")

}


func GetMgo() *mgo.Session {

	if session == nil {
        var err error
        session, err = mgo.Dial("localhost:27017")
        if err != nil {
            panic(err) // no, not really
        }
    }

    return session.Copy()
}

func GetDataBase() *mgo.Database {
    return database
}

func Close(s *mgo.Session) {
    s.Close()
}