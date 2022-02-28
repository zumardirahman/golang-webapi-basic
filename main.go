package main

import (
	"fmt"
	"log"
	"net/http"

	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db Erorr")
	}

	fmt.Println("Database Berhasil Terhubung")

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
