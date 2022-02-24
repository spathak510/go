package main

import (
	"app/db"
	"app/routers"
	"fmt"
	"os"

	_ "github.com/gin-gonic/gin"
	_ "github.com/stretchr/testify/assert"
)

func init() {
	//db.Connect()
	db.ConfigDB()
}

func main() {

	if os.Args != nil && len(os.Args) > 1 {
		fmt.Println(os.Args)
		command := os.Args[1]
		if command == "db" {
			db.Migration()
		} else {
			command = ""
			fmt.Println("No relevant command given!")
		}
		os.Exit(1)
	}

	routers.SetupRouter()
	// r := routers.SetupRouter()

	// port := os.Getenv("port")
	// if len(os.Args) > 1 {
	// 	reqPort := os.Args[1]
	// 	if reqPort != "" {
	// 		port = reqPort
	// 	}
	// }

	// if port == "" {
	// 	port = "80" //localhost
	// }

	// r.Run(":" + port)
}
