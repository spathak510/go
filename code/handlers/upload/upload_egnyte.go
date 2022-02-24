package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

const folderpath  = "storage/egnyte"

func UploadEgnyte(c *gin.Context) {

	file, ferr := c.FormFile("file")
	if ferr != nil {
		fmt.Println(ferr)
		c.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"message": "File not found!",// cast it to string before showing
		})
		return
	}

	path, _ := os.Getwd()
	/*if err := ensureDir("storage"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"message": "Egnyte directory is not found!",
		})
		return
	}*/
	path += "/"+folderpath+"/"+file.Filename
	//fmt.Println("Path1",path)

	c.SaveUploadedFile(file, path)

	requrl := os.Getenv("EGNYTE_UPLOAD_API") + os.Getenv("EGNYTE_UPLOAD_FOLDER") + file.Filename
	//fmt.Println("Path2",requrl)
	data, fExistsErr := os.Open(path)
	if fExistsErr != nil {
		fmt.Println(fExistsErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"code" : http.StatusBadRequest,
			"message": "File not found on path!",
		})
		return
	}
	//fmt.Println(requrl)
	req, _ := http.NewRequest("POST", requrl, data)
	bearer := "Bearer "+os.Getenv("EGNYTE_AUTH_KEY")
	req.Header.Add("Authorization", bearer)
	req.Header.Add("content-type", "text/plain")
	req.Header.Add("name", "file")
	//req.Header.Add("filename", "~/Desktop/test1111.txt")
	req.Header.Add("Content-Disposition", "form-data")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//fmt.Println("amit:")
	//body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	//fmt.Println("amit:",res.StatusCode)
	status := "Failed"
	response := "Failed to upload!"
	if res.StatusCode == 200 {
		fmt.Errorf("bad status: %s", res.Status)
		status = "Success"
		response = "Successfully uploded!"
	}


	c.JSON(http.StatusCreated, gin.H{
		"Code" : http.StatusOK,
		"Status": status,
		"data": response,
	})

}

func ensureDir(dirName string) error {

	err := os.Mkdir(dirName, 0777)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

