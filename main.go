package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int	`json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Meaning", Author: "Viktor Frankle", Quantity: 2},
	{ID: "2", Title: "Rilegion; Unnecessary?", Author: "Sakif Abdullah", Quantity: 10},
	{ID: "3", Title: "Know your mind as Muslim", Author: "Sakif Abdullah", Quantity: 20},
}

func getBooks(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(id string) (*book, error){
	for i, book := range books{
		if book.ID == id{
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

func getBookById(c *gin.Context){
	id := c.Param("id") // "/books/:id"
	book, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg":"Book not found!"})
		return 
	}

	c.IndentedJSON(http.StatusOK, book)
}

func createBooks(c *gin.Context)  {
	var newBook book
	err := c.BindJSON(&newBook)
	if err != nil {
		return 
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func checkoutBook(c *gin.Context)  {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg":"missing id query parameter"})
		return
	}

	book, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg":"book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg":"book not available"})
		return
	}

	book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, book)
}

func returnBook(c *gin.Context)  {
	id, ok := c.GetQuery("id")

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg":"missing id query parameter"})
		return
	}

	book, err := bookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg":"book not found"})
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg":"book not available"})
		return
	}

	book.Quantity += 1
	c.IndentedJSON(http.StatusOK, book)
}

func main()  {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBooks)
	router.GET("/books/:id", getBookById)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}