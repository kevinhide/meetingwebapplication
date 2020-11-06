package daos

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

var mongo_Url string = "mongodb://localhost:27017"
var mongo_DB string = "helpnow"

// GetDB :
func GetDB() *mgo.Database {
	s, err := mgo.Dial(mongo_Url)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	fmt.Println("Connected to", mongo_Url)
	defer s.Close()
	databaseName := mongo_DB
	if len(databaseName) == 0 {
		databaseName = mongo_DB
	}
	return s.Copy().DB(databaseName)
}
