package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Connect(c *gin.Context) {
	// s := db.Session.Clone()

	// defer func() {
	// 	//fmt.Println("calling defer",db.Mongo.Databas)
	// 	s.Close()
	// }()

	// c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}
