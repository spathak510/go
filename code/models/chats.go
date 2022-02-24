package models

import "gopkg.in/mgo.v2/bson"

const (
	CollectionChats = "chats"
)

type Chats struct {
	//_Id 	  	bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Id  		string `json:"id"`
	Datetime 	string `json:"date"`
	Metadata    interface{} `json:"metadata"`
}

type ChatStat struct {
	_Id 	  	bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Id  		string `json:"id"`
	Datetime 	string `json:"datetime"`
}

type RequestId struct {
	ReqId  		string `json:"id"`
}