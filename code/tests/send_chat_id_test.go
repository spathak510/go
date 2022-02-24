package tests

import (
	"app/middlewares"
	"app/handlers/chats"
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)


const Expected_StatusCode  = 201

type TestCase struct {
	requestBody        string
	expectedStatusCode int
}

func TestSendChatId(t *testing.T) {
	tests := []TestCase{
		{`{}`, Expected_StatusCode},
		{`{"id":""}`, Expected_StatusCode},
		{`{"id":"3"}`, Expected_StatusCode},
		{`{"id":"5a5dad9ef35df0e64e190c51"}`, Expected_StatusCode},
	}

	for _, testCase := range tests {
		router := gin.New()
		router.Use(middlewares.Connect)
		router.POST("/send_chat_id", chats.SendChatId)
		json := []byte(testCase.requestBody)
		req, _ := http.NewRequest("POST", "/send_chat_id", bytes.NewBuffer(json))
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		assert.Equal(t, testCase.expectedStatusCode, resp.Code, resp.Body.String())
		if resp.Code == testCase.expectedStatusCode {
			t.Errorf("Passed Case:\nRequest body : %s \n expectedStatus : %d \n observedStatusCode : %d \n responseBody : %s \n", testCase.requestBody, testCase.expectedStatusCode, resp.Code, resp.Body.String())
		} else {
			t.Errorf("Failed Case:\nRequest body : %s \n expectedStatus : %d \n observedStatusCode : %d \n responseBody : %s \n", testCase.requestBody, testCase.expectedStatusCode, resp.Code, resp.Body.String())
		}
	}
}
