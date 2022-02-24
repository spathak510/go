package models

import (
	//"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
)

const (
	CollectionChatDetails = "chatdetails"
)

type Metadata1 struct {
	_Id 	  	bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Id  string `json:"metadata"`
}

type MetadataSet struct {
	Id  string
	Key  string
	Value  interface{}

}

func GetDetails(c *gin.Context, param  map[string]interface{}) []MetadataSet {
	var conditions = bson.M{}
	if len(param) > 0 {
		for key, value := range param {
			fmt.Println(key, value)
			conditions[key]  = value
			//conditions["age"] = bson.M{"$gte": 15}
		}
	}
	if(len(conditions) == 0) {
		conditions = nil
	}

	db := c.MustGet("db").(*mgo.Database)
	data := []MetadataSet{}
	//
	err := db.C(CollectionChatDetails).Find(conditions).All(&data)
	if err != nil {
		c.Error(err)
	}

	return data
}

type ChatDetails struct {
	Id  string
	Key  string
	Value  interface{}

}