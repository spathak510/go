package chats

import (
	"app/models"
	_"fmt"

	//"encoding/json"

	//"bytes"
	//"encoding/json"
	//"fmt"
	"github.com/gin-gonic/gin"
	_ "fmt"
	"net/http"
)

func ChatDetails(c *gin.Context) {

	Reqdata := models.RequestId{}
	c.BindJSON(&Reqdata)
	if len(Reqdata.ReqId) != 24 {
			c.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"message": "Id is missing or incorrect!",// cast it to string before showing
		})
		return
	}
	param := make(map[string]interface{})
	param["id"] = Reqdata.ReqId
	data := models.GetDetails(c,param)

	status := "Failed"
	code := 401
	response := []models.MetadataSet{}
	if(len(data) > 0){
		status = "Success"
		code = 200
		response = data
	}

	c.JSON(http.StatusOK, gin.H{
		"Code" : code,
		"Status": status,
		"Data": response,
	})
}



