package tests

import (
	"app/db"
	"app/middlewares"
	_"app/models"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/stretchr/testify/assert"
	"app/handlers/chats"
	"bytes"
	// "gopkg.in/mgo.v2"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/joho/godotenv"
	//"fmt"
)

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Error loading .env file")
	}
	db.Connect()
}

const ExpectedStatusCode  = 201

type TestStruct struct {
	requestBody        string
	expectedStatusCode int
	//responseBody       string
	//observedStatusCode int
}

func TestCreate(t *testing.T) {
	tests := []TestStruct{
		{`{}`, ExpectedStatusCode},
		{`{"id":""}`, ExpectedStatusCode},
		{`{"id":"3"}`, ExpectedStatusCode},
	}

	for _, testCase := range tests {
		router := gin.New()
		router.Use(middlewares.Connect)
		router.POST("/testCreate", chats.Create)
		json := []byte(testCase.requestBody)//(`{"id":"7878787878"}`)
		req, _ := http.NewRequest("POST", "/testCreate", bytes.NewBuffer(json))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		//assert.Equal(t, testCase.expectedStatusCode, resp.Code, resp.Body.String())
		if resp.Code == testCase.expectedStatusCode {
			t.Errorf("Passed Case:\nRequest body : %s \n expectedStatus : %d \n observedStatusCode : %d \n responseBody : %s \n", testCase.requestBody, testCase.expectedStatusCode, resp.Code, resp.Body.String())
		} else {
			t.Errorf("Failed Case:\nRequest body : %s \n expectedStatus : %d \n observedStatusCode : %d \n responseBody : %s \n", testCase.requestBody, testCase.expectedStatusCode, resp.Code, resp.Body.String())
		}
	}
}



