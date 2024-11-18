package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.String(200, "Welcome to my Page, To visit APIs, please visit /name, /hobby, /dream.")
	})

	router.GET("/name", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.String(200, "Elizabet Yonas Regalzi")
	})
	router.GET("/hobby", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hobby 1": "Codding",
			"hobby 2": "Reading Books",
			"hobby 3": "Listening to Musics",
		})
	})
	router.GET("/dream", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/plain")
		c.String(200, "When you think you have done enough, you can always do more to become best version of yourself.")
	})

	router.ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
