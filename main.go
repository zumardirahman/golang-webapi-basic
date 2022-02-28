package main

import (
	"log"

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

	db.AutoMigrate(&book.Book{})

	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)

	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/books", bookHandler.GetBooks)

	v1.POST("/books", bookHandler.PostBooksHandler)
	//alur post data
	//ke PostBooksHandler (/handler/book.go/PostBooksHandler)

	router.Run(":8888")
}

//Layer GO
//
//1. Main
//2. Handler
//3. Service
//4. Repository
//5. DB
//6. MySql
