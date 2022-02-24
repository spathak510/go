package services

import (
	"crypto/rand"
	"fmt"
	//"github.com/gin-gonic/gin"
	//"net/http"
	"strings"
	"os"
)

//Common functions
func RandomString(str string) string {
	b := make([]byte, 3)
	rand.Read(b)
	id := fmt.Sprintf(str+"%x", b)
	return strings.ToUpper(id)
}

func ConsumeAPI(ParamList map[interface{}]interface{}) string {
	//fmt.Println(len(ParamList["Header"].(interface{})))
	if ParamList["method"] == ""{
		API_Method := "GET"
		_ = API_Method
	}

	if ParamList["url"] == ""{
		return "API Url is missing!"
	} else {
		API_Url := ParamList["url"]
		_ = API_Url
	}

	//if len(ParamList["Header"].(interface{})) > 0{
	//	fmt.Println(len(ParamList["Header"].(interface{})))
	//}
	os.Exit(10)
	return 	"test"
}