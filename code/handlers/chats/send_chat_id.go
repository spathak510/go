package chats

import (
	"app/models"
	"bytes"
	"encoding/json"
	"fmt"
	_ "fmt"

	"github.com/gin-gonic/gin"

	//"io/ioutil"
	"net/http"
)

func TestJWT(c *gin.Context) {
	testData := map[string]string{"test_jwt": "This is testing api for JWT!"}
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    "200",
		"message": "Success",
		"data":    testData,
	})
}
func SendChatId(c *gin.Context) {

	Reqdata := models.RequestId{}
	c.BindJSON(&Reqdata)
	if len(Reqdata.ReqId) != 24 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Id is missing or incorrect!", // cast it to string before showing
		})
		return
	}
	url := "https://api.sparkpost.com/api/v1/transmissions"
	from := "sandbox@sparkpostbox.com"
	subject := "Chat Id"
	mail_body := "Chat id : " + Reqdata.ReqId
	to := "amit.upadhyay@shiftpixy.com"
	parameter := map[string]interface{}{
		"options": map[string]bool{"sandbox": true},
		"content": map[string]string{
			"from":    from,
			"subject": subject,
			"text":    mail_body,
		},
		"recipients": []map[string]string{
			{"address": to},
		},
	}

	bytesRepresentation, err := json.Marshal(parameter)
	_ = bytesRepresentation
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"code":    http.StatusForbidden,
			"message": "Invalid Json Formate!",
			"data":    "",
		})
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Authorization", "bb8725ae7b3a5b94ed33ecfc2eb7ee47e0fc8cef")
	req.Header.Set("Content-Type", "application/json")

	clientReq := &http.Client{}
	resp, err := clientReq.Do(req)
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"code":    http.StatusForbidden,
			"message": "No response from API!",
			"data":    "",
		})
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	status := "Failed"
	response := "message sending error!"
	if resp.Status == "200" {
		status = "Success"
		response = "Message sent to " + to + "!"
	}

	c.JSON(http.StatusCreated, gin.H{
		"Code":   resp.Status,
		"Status": status,
		"data":   response,
	})
}
