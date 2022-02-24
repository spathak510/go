package db

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"gopkg.in/mgo.v2"
)

// var (
// 	Session *mgo.Session

// 	Mongo *mgo.DialInfo
// )
var Dbcon *mongo.Database

func ConfigDB() {
	finalDbUri := ""
	uri := ""
	dbtype := "mongodb://"
	if len(os.Getenv("DB_CONTEXT")) > 0 && (strings.ToLower(os.Getenv("DB_CONTEXT")) == "stg" || strings.ToLower(os.Getenv("DB_CONTEXT")) == "prod") {
		dbtype = "mongodb+srv://"
	}
	if len(os.Getenv("DB_USER")) > 0 {
		uri += os.Getenv("DB_USER") + ":"
	}

	if len(os.Getenv("DB_PASS")) > 0 {
		uri += os.Getenv("DB_PASS") + "@"
	}

	finalDbUri = dbtype + uri + os.Getenv("DB_URL") + "/" + os.Getenv("DB_NAME")
	// fmt.Println(finalDbUri)

	if strings.ToLower(os.Getenv("DB_DEBUG")) == "true" {
		fmt.Println(finalDbUri)
	}

	path, _ := os.Getwd()
	d1 := []byte(finalDbUri)
	err := ioutil.WriteFile(path+"/dblog", d1, 0644)
	check(err)

	if len(finalDbUri) == 0 {
		fmt.Println("DB Uri is not found!")
		os.Exit(2)
	}
	ctx := context.Background()

	// create a mongo client
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(finalDbUri),
	)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected Db!")
	}

	// disconnects from mongo
	//defer client.Disconnect(ctx)
	Dbcon = client.Database(os.Getenv("DB_NAME"))
}

// func Connect() {
// 	finalDbUri := ""
// 	uri := ""
// 	dbtype := "mongodb://"
// 	if len(os.Getenv("DB_CONTEXT")) > 0 && (strings.ToLower(os.Getenv("DB_CONTEXT")) == "stg" || strings.ToLower(os.Getenv("DB_CONTEXT")) == "prod") {
// 		dbtype = "mongodb+srv://"
// 	}
// 	if len(os.Getenv("DB_USER")) > 0 {
// 		uri += os.Getenv("DB_USER") + ":"
// 	}

// 	if len(os.Getenv("DB_PASS")) > 0 {
// 		uri += os.Getenv("DB_PASS") + "@"
// 	}

// 	finalDbUri = dbtype + uri + os.Getenv("DB_URL") + "/" + os.Getenv("DB_NAME")
// 	//fmt.Println(finalDbUri)

// 	if strings.ToLower(os.Getenv("DB_DEBUG")) == "true" {
// 		fmt.Println(finalDbUri)
// 	}

// 	path, _ := os.Getwd()
// 	d1 := []byte(finalDbUri)
// 	err := ioutil.WriteFile(path+"/dblog", d1, 0644)
// 	check(err)

// 	if len(finalDbUri) == 0 {
// 		fmt.Println("DB Uri is not found!")
// 		os.Exit(2)
// 	}

// 	mongo, err := mgo.ParseURL(finalDbUri)
// 	s, err := mgo.Dial(finalDbUri)
// 	if err != nil {
// 		fmt.Printf("Can't connect to mongo, go error %v\n", err)
// 		panic(err.Error())
// 	}
// 	s.SetSafe(&mgo.Safe{})
// 	// fmt.Println("Connected to", finalDbUri)
// 	Session = s
// 	Mongo = mongo

// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}
