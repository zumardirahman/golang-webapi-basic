package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"pustaka-api/book"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) { //public diawali huruf capital agar bisa dipanggil luar paket handler
	c.JSON(http.StatusOK, gin.H{
		"nama": "Zumardi Rahman",
		"bio":  "A Software Engineer",
	})
}

func (h *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title":    "Hello World",
		"subtitle": "My Golang Basic",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	id := c.Param("id") //parameter
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title") //query string
	price := c.Query("price")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func (h *bookHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		//log.Fatal(err) //serber mati

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) { //menampilkan erorr validation
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

			c.JSON(http.StatusBadRequest, gin.H{
				"errors": errorMessages,
			})
			return
		}

	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
