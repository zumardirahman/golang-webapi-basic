package main

import (
	"fmt"
	"log"
	"net/http"

	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Db Erorr")
	}

	//fmt.Println("Database Berhasil Terhubung")
	db.AutoMigrate(&book.Book{})
	//CRUD

	// book := book.Book{}
	// book.Title = "Cari Pengalaman"
	// book.Price = 90000
	// book.Discount = 10
	// book.Rating = 5
	// book.Description = "ini adalah buku yang sangat bagus dari zumardi rahman"

	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("================================")
	// 	fmt.Println("Error vreating book record")
	// 	fmt.Println("================================")
	// }

	//	var book book.Book
	var books []book.Book

	//err = db.Debug().First(&book,2).Error //mengambil primary key
	// err = db.Debug().Find(&books).Error //mengambil bnyak data
	err = db.Debug().Where("rating = ?", 5).Find(&books).Error //mengambil wehere

	if err != nil {
		fmt.Println("================================")
		fmt.Println("Error Finding book record")
		fmt.Println("================================")
	}

	for _, b := range books {

		fmt.Println("title:", b.Title)
		fmt.Println("book object", b)
	}

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
