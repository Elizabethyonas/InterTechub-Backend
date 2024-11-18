package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.String(200, "Welcome to my Golang API, please visit /name, /hobby, /dream.")
}

func NameHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.String(200, "Elizabet Yonas Regalzi")
}

func HobbyHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"hobby 1": "Coding",
		"hobby 2": "Reading Books",
		"hobby 3": "Listening to Music",
	})
}

func DreamHandler(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.String(200, "When you think you have done enough, you can always do more to become the best version of yourself.")
}

func main() {
	r := gin.Default()
	r.GET("/", Handler)
	r.GET("/name", NameHandler)
	r.GET("/hobby", HobbyHandler)
	r.GET("/dream", DreamHandler)

	r.Run(":8080")
}
