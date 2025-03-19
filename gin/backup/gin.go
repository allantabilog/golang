package example

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

var message = "pong"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Message: %s %s %s %s", message, message, message, message)
	})
	return r
}

func postUser(r *gin.Engine) *gin.Engine {
	r.POST("/user/add", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		c.JSON(200, user)
	})
	return r
}

// func oldMain() {
// 	r := setupRouter()
// 	r = postUser(r)
// 	r.Run(":8088")
// }
