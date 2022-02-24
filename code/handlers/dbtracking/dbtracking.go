package dbtracking

import (
	"app/models"
	//"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strings"

	//"log"
	"net/http"
	"reflect"
	//"strings"
)

func ListColle(c *gin.Context) {

	db := c.MustGet("db").(*mgo.Database)
	//Get Collection Names
	CollectionNames, err := db.CollectionNames()
	if err != nil {
		log.Printf("Failed to get coll names: %v", err)
		return
	}
	colMap := make(map[string]string)
	for _, v := range CollectionNames {
		colMap[v] = v
	}

	fmt.Println(colMap)
	fmt.Println(reflect.TypeOf(CollectionNames))
	fmt.Println(CollectionNames)

	m := make(map[string]interface{})
	m["chats"] = models.Chats{}
	m["chatdetails"] = models.ChatDetails{}
	for a,s := range m {
		fmt.Println(a,s)
		b := fmt.Sprintf("%T",s)
		e := strings.Split(b, ".")
		d := e[len(e)-1]
		fmt.Println(d)
		structName := strings.ToLower(d)
		if val, ok := colMap[structName]; ok {
			fmt.Println("Exists" ,val )
			GetFieldNames(c, structName)
		} else {

		}
	}


	//db := c.MustGet("db").(*mgo.Database)
	////Get Collection Names
	//CollectionNames, err := db.CollectionNames()
	//if err != nil {
	//	log.Printf("Failed to get coll names: %v", err)
	//	return
	//}
	//for _, collname := range CollectionNames {
	//	//fmt.Println(collname)
	//	//a := models.collname{}
	//	//var p string = "models.Chats"
	//	//xxx := reflect.ValueOf(collname).String()
	//	//a := &models.CollectionNames[as]{}
	//	//val := reflect.Indirect(reflect.ValueOf(a))
	//	//fmt.Println(val.Type().Field(0).Name)
	//	//
	//	//fmt.Println(reflect.TypeOf(collname))
	//	//if name == "collectionToCheck" {
	//	//	log.Printf("The collection exists!")
	//	//	break
	//
	//	fieldList := db.C(collname)
	//	pipe := fieldList.Pipe(
	//		[]bson.M{
	//			bson.M{
	//				"$project": bson.M{
	//					"arrayofkeyvalue": bson.M{"$objectToArray": "$$ROOT"},
	//				},
	//			},
	//			bson.M{
	//				"$unwind": "$arrayofkeyvalue",
	//			},
	//			bson.M{
	//				"$group": bson.M{
	//					"_id": -1,
	//					"allkeys": bson.M{"$addToSet": "$arrayofkeyvalue.k"},
	//				},
	//			},
	//			bson.M{
	//				"$project": bson.M{
	//					"fieldNames": "$allkeys",
	//				},
	//			},
	//		},
	//	)
	//	result := []bson.M{}
	//	err := pipe.All(&result)
	//	_ = err
	//	fmt.Printf("%+v afwwg", result[0]["fieldName"])
	//	fieldNames := result[0]["fieldNames"].([]interface{})
	//
	//	fmt.Println("/n asfd",fieldNames)
	//	for index, res := range fieldNames{
	//		_=res
	//		val := reflect.Indirect(reflect.ValueOf(&models.MetadataSet{}))
	//
	//		if index == (len(fieldNames) - 1 ) {
	//			break
	//		}
	//
	//		cp2 := strings.ToLower(val.Type().Field(index).Name)
	//
	//		_=cp2
	//
	//		// if cp2 == res[index] {
	//
	//		// }
	//
	//		fmt.Println(index)
	//
	//		fmt.Println(val.Type().Field(index).Name)
	//
	//
	//
	//
	//		//if name == "collectionToCheck" {
	//		//	log.Printf("The collection exists!")
	//		//	break
	//		//}
	//	}
	//	break
	//}
	//var a = &models.collname{}
	//val := reflect.Indirect(reflect.ValueOf(a))
	//fmt.Println(val.Type().Field(0).Name)
	//names, err := db.CollectionNames()
	//if err != nil {
	//	// Handle error
	//	log.Printf("Failed to get coll names: %v", err)
	//	return
	//}
	//
	//// Simply search in the names slice, e.g.
	//for _, name := range names {
	//	fmt.Println(name)
	//	if name == "collectionToCheck" {
	//		log.Printf("The collection exists!")
	//		break
	//	}
	//}
}

func GetFieldNames(c *gin.Context, collname string) []string {
	db := c.MustGet("db").(*mgo.Database)
		fieldList := db.C(collname)
		pipe := fieldList.Pipe(
			[]bson.M{
				bson.M{
					"$project": bson.M{
						"arrayofkeyvalue": bson.M{"$objectToArray": "$$ROOT"},
					},
				},
				bson.M{
					"$unwind": "$arrayofkeyvalue",
				},
				bson.M{
					"$group": bson.M{
						"_id": -1,
						"allkeys": bson.M{"$addToSet": "$arrayofkeyvalue.k"},
					},
				},
				bson.M{
					"$project": bson.M{
						"fieldNames": "$allkeys",
					},
				},
			},
		)
		result := []bson.M{}
		err := pipe.All(&result)
		_ = err
		fmt.Println(result)
		fmt.Println(reflect.TypeOf(result))
		//rawjson, err := json.Marshal(result)
		//fmt.Println(rawjson)
		//ms : make(map[string]interface{})
		//data, err := bson.Marshal(result[0]["fieldName"])
		//fmt.Println(data)
		//fmt.Printf("%+v afwwg", result[0]["fieldName"])
		fmt.Println("==============")
		fieldNames := result[0]["fieldNames"].([]interface{})
		fmt.Println("==============")
		fmt.Printf("/n asfd %=v",fieldNames)
		fmt.Println(reflect.TypeOf(fieldNames))


		//list := make(map[string]interface{})

		for k, v := range result {
			fmt.Println(k,v)
		}


	return []string{"test"}
}

func Createqq(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	thing := models.Thing{}
	err := c.BindJSON(&thing)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionStuff).Insert(thing)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusCreated, thing)
}

func GetOne(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	thing := models.Thing{}
	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionStuff).FindId(oID).One(&thing)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, thing)
}

func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	oID := bson.ObjectIdHex(c.Param("_id"))
	err := db.C(models.CollectionStuff).RemoveId(oID)
	if err != nil {
		c.Error(err)
	}

	// What to do here if this is close to REST
	//c.Redirect(http.StatusMovedPermanently, "/stuff")
	c.Data(204, "application/json", make([]byte, 0))
}

func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	thing := models.Thing{}
	err := c.Bind(&thing)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{ "_id": bson.ObjectIdHex(c.Param("_id")) }
	doc := bson.M{
		"name":		thing.Name,
		"value":	thing.Value,
	}
	err = db.C(models.CollectionStuff).Update(query, doc)
	if err != nil {
		c.Error(err)
	}

	c.Data(http.StatusOK, "application/json", make([]byte, 0))
	//c.Rediret(http.StatusMovedPermanently, "/stuff"
}
