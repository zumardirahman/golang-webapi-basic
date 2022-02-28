package main

import (
	"net/http"

	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/hello", handler.HelloHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler) // :id akan dapat berubah
	v1.GET("/query", handler.QueryHandler)            // ex ?id=232
	v1.POST("/books", handler.PostBooksHandler)

	v2 := router.Group("/v2")
	v2.GET("/", v2RootHandler)

	router.Run(":8888")
}

func v2RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nama": "Zumardi Rahman ver 2",
		"bio":  "A Software Engineer Ver 2",
	})
}
