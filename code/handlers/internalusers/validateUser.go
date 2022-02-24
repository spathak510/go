package internalusers

import (
	"app/db"
	"context"
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const COLLECTION_NAME = "internalUsers"

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func ValidateUser(c *gin.Context) {

	requestData := login{}
	c.BindJSON(&requestData)
	if len(requestData.Username) == 0 || len(requestData.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "401",
			"message": "Username or password is missing!",
		})
		return
	}

	var code string = "401"
	var message string = "Failed"
	var resp string = "Unauthorised User!"

	fieldMap := getOne(requestData.Username, COLLECTION_NAME)

	pwdStr := []byte(requestData.Password)
	pwdMd5 := fmt.Sprintf("%x", md5.Sum(pwdStr))

	if fieldMap["password"] == pwdMd5 {
		code = "200"
		message = "Success"
		resp = "Authorised User!"
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"message": message,
		"data":    resp,
	})
}

func getOne(filterSet string, collection string) map[string]string {
	fieldMap := make(map[string]string)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	dbs := db.Dbcon
	fieldList := dbs.Collection(collection)
	filter := bson.M{"username": bson.M{"$eq": filterSet}}
	tempMap := make(map[string]interface{})
	if err := fieldList.FindOne(ctx, filter).Decode(&tempMap); err == nil {
		if len(tempMap) != 0 {
			for key, value := range tempMap {
				if key != "_id" {
					fieldMap[key] = value.(string)
				}
			}
		}
	}
	return fieldMap
}

func CreateRec(c *gin.Context) {
	requestData := login{}
	c.BindJSON(&requestData)
	if len(requestData.Username) == 0 || len(requestData.Password) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "401",
			"message": "Username or password is missing!",
		})
		return
	}

	var code string = "401"
	var message string = "Failed"
	var resp string = "Failed to add new record!"

	pwdStr := []byte(requestData.Password)
	pwdMd5 := fmt.Sprintf("%x", md5.Sum(pwdStr))

	fieldMap := getOne(requestData.Username, COLLECTION_NAME)
	if len(fieldMap) > 0 {
		resp = "Username is already exists!"
	} else {
		dbs := db.Dbcon
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		res, err := dbs.Collection(COLLECTION_NAME).InsertOne(ctx, bson.D{
			{"username", requestData.Username},
			{"password", pwdMd5},
		})
		if err != nil {
			fmt.Println("", fmt.Errorf("createTask: task for to-do list couldn't be created: %v", err))
		}

		lastid := fmt.Sprintln(res.InsertedID.(primitive.ObjectID).Hex())
		fmt.Println(lastid)
		if len(lastid) > 0 {
			code = "200"
			message = "Success"
			resp = "Successfully added new record!"
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"code":    code,
		"message": message,
		"data":    resp,
	})
}
