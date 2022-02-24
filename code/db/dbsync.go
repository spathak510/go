package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	// "gopkg.in/mgo.v2/bson"

	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	//"gopkg.in/mgo.v2/bson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Migration() {
	fmt.Println("********************************Migration Start*******************************")
	structList := StructList()
	for key, val := range structList {
		_ = val
		fmt.Println("************************Collection:", key, "***************************************")
		isCollectionExists := CollectionExists(key) //Check if collection exists
		// _ = isCollectionExists
		if isCollectionExists == true {
			fmt.Println(key, "Collection exists!")
			resp := SchemaExists(key, val)
			fmt.Println(resp)
		} else {
			fmt.Println(key, "Collection does not exist!")
			structFields := GetStructFields(key, val)
			creatCollection := UpsertFields(key, structFields)
			fmt.Println(creatCollection)
		}
	}
	fmt.Println("********************************Migration end*******************************")
	os.Exit(0)
}

func CollectionExists(collectionName string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	c := Dbcon
	test, err := c.Collection(collectionName).Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
		return false
	}
	count := 0
	for test.Next(context.TODO()) {
		count++
	}
	// fmt.Println(count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

func SchemaExists(collectionName string, schema interface{}) string {
	fmt.Println("Fetching collection fields...")
	fieldMap := GetCollectionFields(collectionName, schema)

	fmt.Println("Fetching structure fields...")
	structFields := GetStructFields(collectionName, schema)
	missingFields := make(map[string]string)
	fmt.Println("Comparing...")
	for k, v := range structFields {
		if kv, ok := fieldMap[k]; ok {
			_ = kv
		} else {
			missingFields[v] = v
		}
	}
	var resp string
	if len(missingFields) > 0 {
		fmt.Println("Missing field found!")
		// fmt.Println(missingFields)
		UpsertFields(collectionName, missingFields)
		resp = "Missing field updated!"
	} else {
		resp = "No change found!"
	}
	return resp
}

func UpsertFields(collectionName string, fields map[string]string) string {
	var addFields = bson.M{}
	if len(fields) > 0 {
		for key, value := range fields {
			addFields[key] = value
		}
	} else {
		addFields = nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	filter := bson.M{"status": "sample"}
	update := bson.M{"$set": addFields}
	col := Dbcon.Collection(collectionName)
	res, err := col.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	_ = err
	_ = res
	return "Added successfully!"

}

func GetStructFields(collectionName string, schema interface{}) map[string]string {
	val := reflect.ValueOf(schema)
	// if its a pointer, resolve its value
	if val.Kind() == reflect.Ptr {
		val = reflect.Indirect(val)
	}

	// should double check we now have a struct (could still be anything)
	if val.Kind() != reflect.Struct {
		log.Fatal("unexpected type")
	}
	structType := val.Type()
	//tableName := structType.Name()
	//fmt.Println("pppppppp",tableName)

	missingFields := make(map[string]string)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldName := strings.ToLower(field.Name)
		missingFields[fieldName] = fieldName
	}
	return missingFields
}

func GetCollectionFields(collectionName string, schema interface{}) map[string]string {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	fieldList := Dbcon.Collection(collectionName)
	filter := bson.M{"status": bson.M{"$eq": "sample"}}

	tempMap := make(map[string]interface{})
	if err := fieldList.FindOne(ctx, filter).Decode(&tempMap); err != nil {
		log.Fatal(err)
	}

	if len(tempMap) == 0 {
		fmt.Println("No fields found in collection")
		structFields := GetStructFields(collectionName, schema)
		creatCollection := UpsertFields(collectionName, structFields)
		_ = creatCollection
		status := make(map[string]string)
		status["status"] = "Fields added to collection " + collectionName
		return status
	}

	fieldMap := make(map[string]string)
	for key, value := range tempMap {
		_ = value
		if key != "_id" {
			fieldMap[key] = key
		}

	}
	return fieldMap
}
