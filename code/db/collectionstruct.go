package db

import (
	_ "fmt"
)

func StructList() map[string]interface{} {
	var structVar = make(map[string]interface{})
	structVar["chats"] = Chats{}
	//structVar["chatdetails"]= ChatDetails{}
	return structVar
}

type Chats struct {
	//_Id 	  	bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Id       string      `json:"id"`
	Datetime string      `json:"date"`
	Metadata interface{} `json:"metadata"`
	Age      string      `json:"age"`
}

type ChatDetails struct {
	Id    string
	Key   string
	Value interface{}
}
