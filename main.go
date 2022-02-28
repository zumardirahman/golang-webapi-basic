package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"nama": "Zumardi Rahman",
			"bio":  "A Software Engineer",
		})
	})

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":    "Hello World",
			"subtitle": "My Golang Basic",
		})
	})

	router.Run(":8888")
}
